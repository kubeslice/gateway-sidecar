@Library('jenkins-library@main') _
dockerImagePipeline(
  script: this,
  service: 'gw-sidecar',
  dockerfile: 'Dockerfile',
  buildContext: '.',
  buildArguments: [PLATFORM:"amd64"]
)
