Comparison of AWS Simple Email Service, Resend, and Loops.

Send test emails using each of the services.

## `ses/`
contains email sending scripts using AWS SES. `smtpsend` does so by connecting to an SES SMTP endpoint. `apisend` does so using the AWS SDK for Go (installed as a dependency) to make calls to the SES API. See the successful sends below:
![ses smtpsend](static/sessmtp_test_success.png)
![ses apisend](static/sesapi_test_success.png)

To use `ses/`, create a `ses/.env` environment file with the following variables:
`SMTP_USERNAME`, `SMTP_PASSWORD`, `FROM_EMAIL`, `TO_EMAIL`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION`.

## `resend/`
contains example Go code for sending a test email via Resend. To use, create a `resend/.env` environment file with your `RESEND_API_KEY`.