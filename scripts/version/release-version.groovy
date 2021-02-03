@Library('curbside-jenkins-library@v3.9.0')
import com.curbside.jenkins.pipeline.curbside.pipelines.CurbsidePipeline
def pipeline = new CurbsidePipeline(this)

pipeline.configuration {

  repository     'terraform-impact'
  force_branch   'main'
  slack_mentions '@mazine'

  branch 'main', {
    shell 'release-version', {
      node                    'generic-s1-standard-1'
      // INCREMENT_VERSION comes from the job parameters
      command_line            "./scripts/version/release-version ${this.env.INCREMENT_VERSION}"
      slack_channels          '#cicd-stag1', '#cicd-platform'
    }
  }
}
