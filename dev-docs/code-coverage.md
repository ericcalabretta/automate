# Code Coverage

[the TOC below is auto-generated by the VSCode "Markdown TOC" plugin; do not edit manually]: #

<!-- TOC depthFrom:2 -->

- [Code Coverage Dashboard](#code-coverage-dashboard)
  - [Viewing UI Code Coverage](#viewing-ui-code-coverage)
  - [Viewing Go Code Coverage](#viewing-go-code-coverage)
  - [Updating Code Coverage Dashboard](#updating-code-coverage-dashboard)
    - [Updating UI code coverage files](#updating-ui-code-coverage-files)
    - [Updating Go code coverage files](#updating-go-code-coverage-files)
- [Running Code Coverage Locally](#running-code-coverage-locally)
  - [UI Code Coverage](#ui-code-coverage)
  - [Go Code Coverage](#go-code-coverage)
- [More Information/Discussion](#more-informationdiscussion)

<!-- /TOC -->

_NB: As of 2020-06-03, UI code coverage [here](https://a2-code-coverage.cd.chef.co/automate-ui) is active but Go coverage [here](https://a2-code-coverage.cd.chef.co/) is not.

## Code Coverage Dashboard

Code Coverage for A2 can be viewed at:

<https://a2-code-coverage.cd.chef.co/>

You need to be on VPN to access the page. It is also linked from the dev environment dashboard, <https://a2-dev.cd.chef.co/>.

This page is currently being built out to be easier to view code coverage. It is currently just an s3 bucket hosting code coverage files. For now it has to periodically be manually updated and the code coverage pages have to be hit directly.

### Viewing UI Code Coverage

UI code coverage can be viewed at:

<https://a2-code-coverage.cd.chef.co/automate-ui/current/>

Past UI code coverage results can be found at the parent page, https://a2-code-coverage.cd.chef.co/automate-ui/, along with a chart showing changes in the last 90 days.

This is automatically regenerated with each merged PR.

### Viewing Go Code Coverage

Currently a work in progress. The #ia_code_metrics group is working on aggregating go coverage across the A2 code base and making it viewable from <https://a2-code-coverage.cd.chef.co/>

### Updating Code Coverage Dashboard

<https://a2-code-coverage.cd.chef.co/> is an s3 bucket in the chef-cd aws account.

To access, ensure you have [Okta AWS](https://chefio.atlassian.net/wiki/spaces/ENG/pages/94259373/SOP-001+-+Gaining+access+to+okta+AWS+accounts+via+the+AWS+API) setup.

You will also need access to the chef-cd account. If you don't have access, request it from #helpdesk.

Run locally the code coverage you want to generate (See [Running Code Coverage Locally](#Running-Code-Coverage-Locally)) then use the aws cli tool that Okta AWS put in place to update the code coverage files on the dashboard.

#### Updating UI code coverage files

Upload the files to the s3 bucket with this command:

`aws --profile chef-cd s3 sync <Local Code Coverage Directory> s3://a2-code-coverage.cd.chef.co/automate-ui/current`

That will put the code coverage in the current folder, so it's easy to always access the last time code coverage was run.

Also add the code coverage files into a `YYYY-MM-DD` so a record of code coverage over time is kept.

`aws --profile chef-cd s3 sync <Local Code Coverage Directory> s3://a2-code-coverage.cd.chef.co/automate-ui/YYYY-MM-DD`

You can see previous code coverage files that have been updated by listing the contents of the directory:

`aws --profile chef-cd s3 ls s3://a2-code-coverage.cd.chef.co/automate-ui --recursive`

#### Updating Go code coverage files

Currently a work in progress.

## Running Code Coverage Locally

### UI Code Coverage

In A2, go to the automate-ui directory. From this directory, run `make test`. This will run the unit and e2e tests and output code coverage results to a `coverage` directory. Open the index.html file in this directory to view the report.

How does 'make test' work:
The makefile triggers `npm` and uses the run command; this is just a general runner that is looking for a script to run.
The scripts it uses are defined in the `package.json` file.
There you will find a test script that is just calling `ng test` with a code coverage option.
`ng` is the angular-cli, which uses karma to do testing and provide code coverage. Karma config can be found in `karma.conf.js`.

For further details:

- [Read this page in the wiki](https://chefio.atlassian.net/wiki/spaces/ENG/pages/128581633/Code+Coverage), which gives more detail and shows example output.
- See the [ui-development docs](./ui-development.md) for details on setting up a test runner.

### Go Code Coverage

Currently running the Go unit tests for a component will output the Go code coverage for that component.

The #ia_code_metrics group is working on code to aggregate Go code coverage across all the Go components and make this available via the code coverage dashboard.

## More Information/Discussion

The group at #ia_code_metrics is working on building out this functionality. Feel free to drop in and ask questions there. Release Engineering is also helping as we explore further options and making this easier to use and get it fully integrated into the pipeline.
