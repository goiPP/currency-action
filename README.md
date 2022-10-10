# Currency Action
An action to convert between currencies both physical and digital. You can use
this action and schedule to run per need to automatically get the latest exchange rate.


## Specification
- This action use Forex(FX)'s CURRENCY_EXCHANGE_RATE, please information about api [here](https://rapidapi.com/alphavantage/api/alpha-vantage)
- Currency codes list can be found in: 
  - [Physical currency](https://www.alphavantage.co/physical_currency_list/)
  - [Digital currency](https://www.alphavantage.co/digital_currency_list/)
- Implement in Go 1.19
- Using [Docker container GitHub Actions](https://docs.github.com/en/actions/creating-actions/creating-a-docker-container-action)
- Require to use on linux GitHub Actions runner

## Usage
#### Inputs:
- `to_currency`: a physical currency or digital/crypto currency you would like to convert to. eg: USD, BTC
- `from_currency`: a physical currency or digital/crypto currency you would like to convert from. eg: THB, BTC
- `API_KEY`: request a free API_KEY in [Alpha Vantage API page](https://rapidapi.com/alphavantage/api/alpha-vantage) 
and set it as a repository secret.

```yaml
name: Your Workflow Name
on: [push, workflow_dispatch]

jobs:
  job_name:
    runs-on: ubuntu-latest
    steps:
      - name: Currency Convert
        uses: goiPP/currency-action@master
        env:
          API_KEY: ${{ secrets.API_KEY }}
        with:
          to_currency: THB
          from_currency: EUR
```
Result:
```yaml
1 EUR = 36.76300000 THB
```

#### Reference links, Thank you!:
- [Article about writing GitHub Actions in Go](https://www.sethvargo.com/writing-github-actions-in-go/)
- [GitHub Repo of Unofficial GitHub Actions SDk for Go](https://github.com/sethvargo/go-githubactions)
- [Document of GitHub Actions Go SDK](https://pkg.go.dev/github.com/sethvargo/go-githubactions)
- [Using environment variable in Go](https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66)
