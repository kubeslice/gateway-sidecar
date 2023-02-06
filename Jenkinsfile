@Library('jenkins-library@opensource') _
dockerImagePipeline(
  script: this,
  service: 'gw-sidecar',
  dockerfile: 'Dockerfile',
  buildContext: '.',
  buildArguments: [PLATFORM:"amd64"]
  
)
