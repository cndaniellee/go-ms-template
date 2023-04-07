pipeline {
    agent any

    stages {
        stage('Git Checkout') {
            steps{
                checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[credentialsId: 'VM-Gitlab', url: 'http://192.168.2.220:8200/root/go-ms-template']])
            }
        }
        stage('Docker Build') {
            steps {
                script {
                    docker.withRegistry('https://192.168.2.220:8443', 'VM-Harbor') {
                        def img = docker.build("goms/${MODULE}-${SERVICE}:${VERSION}", "-f service/${MODULE}/${SERVICE}/Dockerfile .")
                        img.push()
                    }
                }
            }
        }
        stage('Image Run') {
            steps {
                script {
                    if (params.DEV_ENV) {
                        try {
                            sh "docker stop goms_service_${MODULE}_${SERVICE}"
                            sh "docker rm goms_service_${MODULE}_${SERVICE}"
                        } catch (e) {
                            echo "${MODULE}-${SERVICE} not running"
                        }
                        sh "docker run -id --name goms_service_${MODULE}_${SERVICE} --network host goms/${MODULE}-${SERVICE}:${VERSION}"
                    } else {
                        sshPublisher(publishers: [sshPublisherDesc(configName: 'k8s-master', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: '', execTimeout: 120000, flatten: true, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: "doc/kube/${MODULE}-${SERVICE}.yaml")], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
                        sshPublisher(publishers: [sshPublisherDesc(configName: 'k8s-master', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: "kubectl apply -f /root/deploy/${MODULE}-${SERVICE}.yaml", execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
                    }
                }
            }
        }
    }
}
