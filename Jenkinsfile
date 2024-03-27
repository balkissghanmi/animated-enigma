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
                sh '/usr/local/go/bin/go build -o goL'
                sh '/usr/local/go/bin/go test'


            }
        }
         stage('Build') {
            steps {
                script {
                    def govulncheck = sh(script: 'go get -u github.com/arthurbailao/govulncheck', returnStdout: true).trim()
                    def goimports = sh(script: 'go get golang.org/x/tools/cmd/goimports', returnStdout: true).trim()
                    def golint = sh(script: 'go get -u golang.org/x/lint/golint', returnStdout: true).trim()
                    def staticcheck = sh(script: 'go get honnef.co/go/tools/cmd/staticcheck', returnStdout: true).trim()
                    def golangciLint = sh(script: 'go install github.com/golangci/golangci-lint/cmd/golangci-lint', returnStdout: true).trim()
                    def gosec = sh(script: 'go get github.com/securego/gosec/cmd/gosec/...', returnStdout: true).trim()

                    sh "echo 'Installed govulncheck: $govulncheck'"
                    sh "echo 'Installed goimports: $goimports'"
                    sh "echo 'Installed golint: $golint'"
                    sh "echo 'Installed staticcheck: $staticcheck'"
                    sh "echo 'Installed golangci-lint: $golangciLint'"
                    sh "echo 'Installed gosec: $gosec'"
                }
            }
        }

        stage('Unit Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Static Analysis') {
            steps {
                sh 'go vet ./...'
                sh 'go tool vet --shadow .'
                sh 'aligncheck ./...' // Assuming aligncheck is installed globally
                sh 'go-critic ./...' // Assuming go-critic is installed globally
                sh 'golint ./...'
                sh 'goimports -w .'
                sh 'gofmt -s -w .'
                sh 'staticcheck ./...'
                sh 'gosec ./...'
                sh 'golangci-lint run'
            }
        }
    
    }
    }