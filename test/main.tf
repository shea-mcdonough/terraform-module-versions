module "A_latest" {
  source = "https://artifacts.dox.support/repository/maven-releases/groupA/artifactID1/1.0.0/artifactID-2.0.0.tgz"
}
module "A_outdated" {
  source = "https://artifacts.dox.support/repository/maven-releases/groupA/artifactID1/1.0.0/artifactID-1.0.0.tgz"
}