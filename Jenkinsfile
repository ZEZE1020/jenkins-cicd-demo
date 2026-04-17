// Declarative Jenkinsfile for a robust CI/CD pipeline
pipeline {
  agent any
  environment {
    DOCKER_IMAGE = "ogembog/jenkins-cicd-demo:${env.BUILD_NUMBER}"
  }
  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }
    stage('Master Agent Example (Demo Only)') {
      when {
        expression { env.RUN_MASTER_AGENT_DEMO == 'true' }
      }
      agent { label 'master || built-in' }
      steps {
        echo 'Demo-only stage: running on Jenkins master/built-in agent. Use dedicated agents in production.'
        sh 'echo "Agent node: $(hostname)"'
      }
    }
    stage('Install Dependencies') {
      steps {
        sh 'go mod download'
      }
    }
    stage('Quality Checks (Parallel)') {
      parallel {
        stage('Lint') {
          steps {
            sh '''
              gofmt -l . | tee lint-report.txt
              if [ -s lint-report.txt ]; then
                echo "Lint failed: run gofmt on the files listed above."
                exit 1
              fi
            '''
          }
        }
        stage('Test') {
          steps {
            sh '''
              set -o pipefail
              go test -v ./... | tee test-output.txt
            '''
          }
        }
      }
    }
    stage('Build') {
      steps {
        sh 'go build -o main ./cmd/server'
        sh 'docker build -t $DOCKER_IMAGE .'
      }
    }
    stage('Scan Image (Trivy)') {
      steps {
        sh '''
          if command -v trivy >/dev/null 2>&1; then
            trivy image --no-progress --severity HIGH,CRITICAL --exit-code 1 "$DOCKER_IMAGE"
          else
            docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
              aquasec/trivy:0.52.2 image --no-progress --severity HIGH,CRITICAL --exit-code 1 "$DOCKER_IMAGE"
          fi
        '''
      }
    }
    stage('Push Image') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-credentials', usernameVariable: 'DOCKERHUB_USER', passwordVariable: 'DOCKERHUB_PASS')]) {
          sh '''
            echo "$DOCKERHUB_PASS" | docker login -u "$DOCKERHUB_USER" --password-stdin
            docker push "$DOCKER_IMAGE"
            docker logout
          '''
        }
      }
    }
    stage('Deploy') {
      steps {
        sh 'chmod +x scripts/deploy.sh'
        sh './scripts/deploy.sh'
      }
    }
    stage('Verify Deployment') {
      steps {
        echo 'Running post-deployment verification...'
        sh 'sleep 10'
        sh 'curl -f http://localhost:3000/health'
        sh 'curl -f http://localhost:3000/dashboard'
      }
    }
  }
  post {
    always {
      archiveArtifacts artifacts: 'test-output.txt,lint-report.txt', allowEmptyArchive: true
    }
  }
}
