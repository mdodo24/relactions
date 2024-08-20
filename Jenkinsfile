pipeline {
    agent any
    
    environment {
        GO_VERSION = '1.20'  // Specify the Go version you're using
        DOCKER_REGISTRY = 'gcr.io/k8s-hands-on-424507'  // GCR registry URL
        IMAGE_NAME = 'relactions-server'  // Your Docker image name
        GCP_SERVICE_ACCOUNT_KEY = credentials('gcr-service-account')  // Jenkins credentials ID for GCR
    }

    stages {
        // stage('Checkout') {
        //     steps {
        //         // Clone the repository
        //         git url: 'https://your-repo-url.git', branch: 'main'
        //     }
        // }

        stage('Install Go') {
            steps {
                // Install Go
                sh "wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"
                sh "sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz"
                sh "export PATH=$PATH:/usr/local/go/bin"
            }
        }

        stage('Build') {
            steps {
                // Build the Go project
                sh 'go mod download' // Install dependencies
                sh 'go build -o relactions .'
            }
        }

        stage('Test') {
            steps {
                // Run tests
                sh 'go test ./...'
            }
        }

        stage('Lint') {
            steps {
                // Run linter
                sh 'go vet ./...'
            }
        }

        stage('Docker Build') {
            steps {
                // Build Docker image
                script {
                    def dockerImage = docker.build("${DOCKER_REGISTRY}/${IMAGE_NAME}:${env.BUILD_NUMBER}")
                }
            }
        }

        stage('Authenticate Docker to GCR') {
            steps {
                script {
                    // Write the service account key to a file
                    def keyFile = "${env.WORKSPACE}/k8s-hands-on-424507-f59c6ef3307e.json"
                    writeFile file: keyFile, text: GCP_SERVICE_ACCOUNT_KEY

                    // Authenticate Docker to GCR using the service account key
                    sh "gcloud auth activate-service-account --key-file=${keyFile}"
                    sh "gcloud auth configure-docker --quiet"
                }
            }
        }

        stage('Docker Push') {
            steps {
                // Push Docker image to GCR
                script {
                    docker.withRegistry("https://${DOCKER_REGISTRY}") {
                        docker.image("${DOCKER_REGISTRY}/${IMAGE_NAME}:${env.BUILD_NUMBER}").push()
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                // Deploy the Docker image, e.g., to Kubernetes or any container service
                sh 'echo "Deploying Docker image..."'
                // Add your deployment steps here
            }
        }
    }

    post {
    always {
        // Clean up the service account key file
        sh "rm -f ${env.WORKSPACE}/gcloud-service-key.json"
    }
}
}
