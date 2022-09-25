pipeline {
    agent any 
    tools {
        go 'go-1.19'
    }
    environment {
        GO119MODULE = 'on'
    }
        stages {
            /*
            stage('Checkout git') {
                steps {
                    git branch: 'main',
                        credentialsId: 'github-ssh',
                        url: 'git@github.com:pajerr/supplement-tracker'
                }
            }*/
            /*
             stage('Run') {
                steps {
                    sh "ls -l"
                    sh "make run"
                }
            }*/           

            stage("unit-test") {
                steps {
                    echo 'UNIT TEST EXECUTION STARTED'
                    //sh "go test -v"
                    sh 'make unit-tests'
                }
            }
            
            stage('Build') {
                steps {
                    sh "make build"
                }
            }
        }

}
