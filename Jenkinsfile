@Library('jenkins-library@master') _
dockerImagePipeline(
  script: this,
  service: 'kubeslice/gw-sidecar',
  dockerfile: 'Dockerfile',
  buildContext: '.',
  buildArguments: [PLATFORM:"amd64"]
)
