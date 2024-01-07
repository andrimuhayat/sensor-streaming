pipeline {
    agent any
    environment {
        GCR_ID = 'leu-indonesia'
        IMAGE_NAME = 'sensor-streaming'
    }
    stages {
        stage("Checkout code") {
            steps {
                checkout scm
            }
        }
        stage("Build image dev") {
            when {
                expression { return env.GIT_BRANCH == "origin/dev" }
            }
            steps {
                script {
                    myapp = docker.build("gcr.io/${env.GCR_ID}/${env.IMAGE_NAME}-dev:${env.BUILD_ID}")
                }
            }
        }
        stage("Build image production") {
            when {
                expression { return env.GIT_BRANCH == "origin/prod" }
            }
            steps {
                script {
                    myapp = docker.build("gcr.io/${env.GCR_ID}/${env.IMAGE_NAME}-prod:${env.BUILD_ID}")
                }
            }
        }
        stage("Push image") {
            steps {
                script {
                    docker.withRegistry('https://gcr.io', 'gcr:leu-gcr') {
                            myapp.push("latest")
                            myapp.push("${env.BUILD_ID}")
                    }
                }
            }
        }
        stage('Deploy to k8s-dev') {
            when {
                expression { return env.GIT_BRANCH == "origin/dev" }
            }
            steps{
                withKubeConfig([credentialsId: 'dev-kubeconfig']) {
                    sh "kubectl apply -f deployment-dev.yaml"
                }
            }
        }

        stage('Deploy to k8s-prod') {
            when {
                expression { return env.GIT_BRANCH == "origin/prod" }
            }
            steps{
                withKubeConfig([credentialsId: 'dev-kubeconfig']) {
                    sh "kubectl apply -f deployment-prod.yaml"
                }
             }
        }

        stage('Remove Unused docker image dev') {
            when {
                expression { return env.GIT_BRANCH == "origin/dev" }
            }
            steps{
                sh "docker rmi gcr.io/${env.GCR_ID}/${env.IMAGE_NAME}-dev:${env.BUILD_ID}"
            }
        }

        stage('Remove Unused docker image prod') {
            when {
                expression { return env.GIT_BRANCH == "origin/prod" }
            }
            steps{
                sh "docker rmi gcr.io/${env.GCR_ID}/${env.IMAGE_NAME}-prod:${env.BUILD_ID}"
            }
        }
    }
}
