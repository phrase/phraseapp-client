node {
  stage "Checkout"
  checkout scm

  stage "Test"
  sh "bash build/test.sh"

  stage "Build"
  sh "bash build/build.sh"
}
