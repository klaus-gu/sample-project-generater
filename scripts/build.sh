#!/bin/bash
rm -rf  /usr/local/go/src/sample-project-generater/$2
mvn archetype:generate -DgroupId=$1 -DartifactId=$2\
    -DarchetypeArtifactId=newsample  -DinteractiveMode=false\
    -DarchetypeCatalog=local -DarchetypeGroupId=com.ovopark.dc\
    -DarchetypeArtifactId=sample-project-archetype

#rm -rf  /usr/local/java/$2
#mv /usr/local/go/src/sample-project-generater/$2  /usr/local/java
