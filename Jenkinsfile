@Library('jenkins-library@opensource-release-multiarch') _
dockerImagePipeline(
  script: this,
  services: ['gw-sidecar'],
  dockerfiles: ['Dockerfile'],
  pushed: true,
  buildArgumentsList: [
    // [ENV: 'production', PLATFORM: 'linux/amd64,linux/arm64']
    [ENV: 'production', PLATFORM: 'linux/amd64']
]
  
)
