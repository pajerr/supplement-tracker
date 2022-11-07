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
