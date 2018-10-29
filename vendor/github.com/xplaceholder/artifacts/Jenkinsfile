pipeline {
    agent { docker { image 'golang:1.10.3' } }
    stages {
        stage('prepare source') {
            steps {
                sh 'echo "hello"'
            }
        }
        stage('golint') {
            steps {
                sh 'golint ./...'
            }
        }
        stage('unit test') {
            steps {
                sh 'ginkgo -r'
            }
        }
    }
}