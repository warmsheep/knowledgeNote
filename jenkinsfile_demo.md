

mvn -f static-code-analysis-example/pom.xml clean package checkstyle:checkstyle findbugs:findbugs cobertura:cobertura pmd:pmd -DskipDockerBuild=true

```
step([$class: 'JUnitResultArchiver', testResults: testResultsPath])
step([$class: 'hudson.plugins.checkstyle.CheckStylePublisher', pattern: '**/target/checkstyle-result.xml', unstableTotalAll:'0',unhealthy:'100',healthy:'100'])
step([$class: 'PmdPublisher', pattern: '**/target/pmd.xml'])
step([$class: 'FindBugsPublisher', pattern: '**/findbugsXml.xml'])       
step([$class: 'CoberturaPublisher', autoUpdateHealth: false, autoUpdateStability: false, coberturaReportFile: 'output/coverage/cobertura-coverage.xml', failUnhealthy: false, failUnstable: false, maxNumberOfBuilds: 0, onlyStable: false, sourceEncoding: 'ASCII', zoomCoverageChart: false])

```         


pom.xml文件加入

<reporting>
      <plugins>
        <plugin>
          <groupId>org.apache.maven.plugins</groupId>
          <artifactId>maven-checkstyle-plugin</artifactId>
          <version>2.17</version>
          <reportSets>
            <reportSet>
              <reports>
                <report>checkstyle</report>
              </reports>
            </reportSet>
          </reportSets>
        </plugin>
        <plugin>
          <groupId>org.apache.maven.plugins</groupId>
          <artifactId>maven-pmd-plugin</artifactId>
          <version>3.7</version>
        </plugin>
        <plugin>
          <groupId>org.codehaus.mojo</groupId>
          <artifactId>findbugs-maven-plugin</artifactId>
          <version>3.0.4</version>
        </plugin>
      </plugins>
</reporting>


node {
   sh 'pwd'
   nodes = sh (script: '../parameters_JenkinsfileParameters@script/list_nodes.sh', returnStdout: true).trim()
}
parameters {
        choice(name: 'Invoke_Parameters', choices:"Yes\nNo", description: "Do you whish to do a dry run to grab parameters?" )
        choice(name: 'Nodes', choices:"${nodes}", description: "")
}


list_nodes.sh 内容
echo "node_one"
echo "node_two"

list_version.sh
echo "v1"
echo "v2"
echo "v3"
