## Setup

# Clone Repo
```
git clone git@github.com:shea-mcdonough/terraform-module-versions.git
cd terraform-module-versions
```

# Setup Nexus
```
docker run -d -p 8081:8081 --name nexus sonatype/nexus3
```
If using dox-cloud, also run:
```
dox-cloud forward 8081
```
Get the admin credentials:
```
docker exec -it nexus bash
cat /nexus-data/admin.password
```
From your terminal, create an empty file to upload
```
tar czvf empty.tgz --files-from=/dev/null
```
From your brower, go to: http://localhost:8081
Click on `Upload` and under `maven-releases` upload `empty.tgz` twice using the following settings:
```
File: empty.tgz
Extension: tgz
Group ID: groupA
Artifact ID: artifactID
Version: 1.0.0 for first upload then 2.0.0 for second upload
```
Run the go module to check for the terraform module versions:
```
go run main.go check test
```
We expect to see that only one module needs to be updated:
```
| UPDATE? |    NAME    | CONSTRAINT | VERSION | LATEST MATCHING | LATEST |
|---------|------------|------------|---------|-----------------|--------|
| (Y)     | A_outdated |            | 1.0.0   |                 | 2.0.0  |
```
