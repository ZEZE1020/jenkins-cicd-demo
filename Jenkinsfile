// Declarative Jenkinsfile for CI/CD pipeline
pipeline {
  agent any
  environment {
    DOCKERHUB_CREDENTIALS = credentials('dockerhub-credentials') // Placeholder
    DOCKER_IMAGE = "ogembog/jenkins-cicd-demo:${env.BUILD_NUMBER}"
  }
  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }
    stage('Install Dependencies') {
      steps {
        sh 'go mod download'
      }
    }
    stage('Run Tests') {
      steps {
        sh 'go test -v ./...'
      }
    }
    stage('Build Application') {
      steps {
        sh 'go build -o main .'
      }
    }
    stage('Build Docker Image') {
      steps {
        sh 'docker build -t $DOCKER_IMAGE .'
      }
    }
    stage('Push Docker Image') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-credentials', usernameVariable: 'DOCKERHUB_USER', passwordVariable: 'DOCKERHUB_PASS')]) {
          sh 'echo $DOCKERHUB_PASS | docker login -u $DOCKERHUB_USER --password-stdin'
          sh 'docker push $DOCKER_IMAGE'
        }
      }
    }
    stage('Deploy') {
      steps {
        echo 'Deploy step: Replace with actual deployment logic.'
      }
    }
  }
}
