pipeline {
  
    agent any
    tools{
        go '1.22.0'
    }
    environment {
    PATH = "/usr/local/go/bin:${env.PATH}"
    DOCKERHUB_USERNAME = 'balkissd'
    STAGING_TAG = "${DOCKERHUB_USERNAME}/go:v1.0.0"
}
    stages {
        stage('Checkout Git') {
            steps {
                script {
                    git branch: 'main',
                        url: 'https://github.com/balkissghanmi/animated-enigma.git',
                        credentialsId: 'test' 
                }
            }
        }
        stage("go version") {
            steps {
                sh 'pwd'
                sh ' go version'
                sh '/usr/local/go/bin/go build -o myapp'
                sh '/usr/local/go/bin/go test'


            }
        }
        stage('Unit Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Static Analysis') {
            steps {
               // sh 'go vet ./...'
               // sh 'go tool vet --shadow .'
               // sh 'aligncheck ./...' // Assuming aligncheck is installed globally
                //sh 'go-critic ./...' // Assuming go-critic is installed globally
                sh'gosec -exclude=fmt,errors ./...'
                sh 'golint ./...'
                //sh 'goimports -w .'
                //sh 'gofmt -s -w .'
                //sh 'staticcheck ./...'
                //sh 'golangci-lint run'
            }
        }
    
    }
    }