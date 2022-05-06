@Library('jenkins-library@opensource') _
dockerImagePipeline(
  script: this,
  service: 'aveshasystems/gw-sidecar',
  dockerfile: 'Dockerfile',
  buildContext: '.',
  buildArguments: [PLATFORM:"amd64"]
)
