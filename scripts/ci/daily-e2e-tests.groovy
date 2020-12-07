@Library('curbside-jenkins-library@v3.8.0')
import com.curbside.jenkins.pipeline.curbside.pipelines.CurbsidePipeline
def pipeline = new CurbsidePipeline(this)

pipeline.configuration {

  cron           'H 12 * * *' // runs once every day
  repository     'terraform-impact'
  force_branch   'main'
  slack_mentions '@vseguin', '@mazine'

  branch 'main', {
    test 'e2e-tests', {
      node                    'generic-s1-standard-1'
      command_line            'make ci-integration-tests'
      fail_when_skipped_tests false
      env_vars                'GITHUB_USERNAME=curbsidebot'
      encrypted_env_vars      'GITHUB_PASSWORD=curbsidebot_github_token'
      result_file             'report.xml'
      slack_channels          '#cicd-platform'
    }
  }
}
