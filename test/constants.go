package test

// SSLTestPath is the path intermediary SSL certificates are created in for test cases
const SSLTestPath = "ssl"

// CACertFolder is the folder at which intermediary CA certificates are created in for test cases
const CACertFolder = "ssl/ca"

// CACertPath is the path of the intermediary CA certificate
const CACertPath = "ssl/ca/ca_cert.pem"

// RegistrationCert is a sample relay proxy registration cert
const RegistrationCert = "-----BEGIN CERTIFICATE-----\nMIICtTCCAjygAwIBAgIUVdro/ObWSYNBWo3CoyHcc4/PtHswCgYIKoZIzj0EAwIw\nbzELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQHDAhFdmFu\nc3RvbjEbMBkGA1UECgwSYmxvWHJvdXRlIExBQlMgSW5jMR0wGwYDVQQDDBRibG9Y\ncm91dGUudGVzdG5ldC5DQTAeFw0yMTEyMDYwODQ1MDlaFw0zMTEyMDQwMDAwMDBa\nMGwxCzAJBgNVBAYTAiAgMQswCQYDVQQIDAIgIDELMAkGA1UEBwwCICAxGzAZBgNV\nBAoMEmJsb1hyb3V0ZSBMQUJTIEluYzEmMCQGA1UEAwwdYmxvWHJvdXRlLnRlc3Ru\nZXQucmVsYXlfcHJveHkwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAARdRXcc5IZ9lKus\nnQh9DWBawedHtpxy8VfztWBvB64oNMMV9H8vb2XRNHhwUMo3SG/SFh+y/fKGaBSQ\nsuZMreLnP4SeQDR1PfibfyhmkwCQyDKQ9jhyKADm0K8X1wBLlQyjgZswgZgwHwYD\nVR0jBBgwFoAUkN1jVx3ISR8ATM29MLuIZ0K0dJEwKAYDVR0RBCEwH4IdYmxvWHJv\ndXRlLnRlc3RuZXQucmVsYXlfcHJveHkwDAYDVR0TAQH/BAIwADAOBgNVHQ8BAf8E\nBAMCBeAwFAYFPoJNolwEC1JFTEFZX1BST1hZMBcGBT6CTaJeBA5ibG9Ycm91dGUg\nTEFCUzAKBggqhkjOPQQDAgNnADBkAjAVVCopO1KqhH4NUlBATOQSgGK0Kp2S2f3O\nmm4UinZcb8gXEQx4x2porpgNsRYtYjsCMAp6DMQxiIuLz1syPEvTByR4wmtNTlRf\nsbfs0ShKmVhQSBIbnIKerpX+Pl1YYoSeqA==\n-----END CERTIFICATE-----\n"

// RegistrationKey is a sample relay proxy registration key
const RegistrationKey = "-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDDAvjme5SkD0OrhdntQurzuvjg8S620n/q2dWEd4qwTyrscWZ5eHgQM\neZPK5gp+wv6gBwYFK4EEACKhZANiAARdRXcc5IZ9lKusnQh9DWBawedHtpxy8Vfz\ntWBvB64oNMMV9H8vb2XRNHhwUMo3SG/SFh+y/fKGaBSQsuZMreLnP4SeQDR1Pfib\nfyhmkwCQyDKQ9jhyKADm0K8X1wBLlQw=\n-----END EC PRIVATE KEY-----\n"

// PrivateCert is a sample relay proxy private cert
const PrivateCert = "-----BEGIN CERTIFICATE-----\nMIICzTCCAlOgAwIBAgIUNmdelk/FEuWx4ZnHSj8DWwslUBowCgYIKoZIzj0EAwIw\nbzELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQHDAhFdmFu\nc3RvbjEbMBkGA1UECgwSYmxvWHJvdXRlIExBQlMgSW5jMR0wGwYDVQQDDBRibG9Y\ncm91dGUudGVzdG5ldC5DQTAeFw0yMzEwMjMyMTIxNTVaFw0zMzEwMTMwMDAwMDBa\nMGwxCzAJBgNVBAYTAiAgMQswCQYDVQQIEwIgIDELMAkGA1UEBxMCICAxGzAZBgNV\nBAoTEmJsb1hyb3V0ZSBMQUJTIEluYzEmMCQGA1UEAwwdYmxvWHJvdXRlLnRlc3Ru\nZXQucmVsYXlfcHJveHkwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAATytPpc10LjWK5X\na+Cd19jNoC2O7K/pU1b8itOF9viDAe3iExyk6ChjMzsUcmj08KrYqAhp1dYK/S5H\nyUcZT89ic+KBrL6UZ4qoGpb51leLGoYjVGxgxHRjmekx8yih+8ejgbIwga8wHwYD\nVR0jBBgwFoAUkN1jVx3ISR8ATM29MLuIZ0K0dJEwDAYDVR0TAQH/BAIwADAOBgNV\nHQ8BAf8EBAMCBeAwFAYFPoJNolwEC1JFTEFZX1BST1hZMC0GBT6CTaJdBCQzZTZj\nMWFiZS1mYzdjLTRlNTMtOWNiZS1hNTUzYjkzMGRkMzAwFwYFPoJNol4EDmJsb1hy\nb3V0ZSBMQUJTMBAGBT6CTaJfBAdnZW5lcmFsMAoGCCqGSM49BAMCA2gAMGUCMQD+\ntUzg7JhIA4ZTZiglRAgOtC8JpzwFYl6oLrUXjKknIvmRpUUzUAO6EZ4fJThAYkoC\nMHqtXZtvFqJUuJObqdxKXqhu0+mkvLJvBSG+29dQP3cJoHQwsS7gVV4vlH7FYg1a\n4Q==\n-----END CERTIFICATE-----\n"

// PrivateKey is a sample relay proxy private key
const PrivateKey = "-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDB5RCTQ8rvYyoixxC/QD9wQpsvhh1XZH1p/1zcc4gbkypEUcbUDHAu2\nvnsHevQZTiSgBwYFK4EEACKhZANiAATytPpc10LjWK5Xa+Cd19jNoC2O7K/pU1b8\nitOF9viDAe3iExyk6ChjMzsUcmj08KrYqAhp1dYK/S5HyUcZT89ic+KBrL6UZ4qo\nGpb51leLGoYjVGxgxHRjmekx8yih+8c=\n-----END EC PRIVATE KEY-----\n"

// CACert is a sample CA certificate
const CACert = "-----BEGIN CERTIFICATE-----\nMIICkDCCAhWgAwIBAgIUQGhfIhpMxaSE0r/jjB35VclkTYgwCgYIKoZIzj0EAwIw\nazELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQHDAhFdmFu\nc3RvbjEbMBkGA1UECgwSYmxvWHJvdXRlIExBQlMgSW5jMRkwFwYDVQQDDBBibG9Y\ncm91dGUuZGV2LkNBMB4XDTIwMTEwOTIyMzY0NloXDTQ4MDMyNzAwMDAwMFowazEL\nMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQHDAhFdmFuc3Rv\nbjEbMBkGA1UECgwSYmxvWHJvdXRlIExBQlMgSW5jMRkwFwYDVQQDDBBibG9Ycm91\ndGUuZGV2LkNBMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEd5IkD4wqWVGbq0jCehjr\nyvOEkbD5vCYIes4UwRH9Z7YSfeKQrOSaW1LUzHCYvqOTLfgHoAC0lS1v9+OlT/LZ\nboey8h4TZwLZ9zLbHzyVcTww01ZFNeBndQ2EmdaSYdqKo3oweDAfBgNVHSMEGDAW\ngBQHnCDGbOz7WL0VB9Kd0YdQSpsnQzAdBgNVHQ4EFgQUB5wgxmzs+1i9FQfSndGH\nUEqbJ0MwGwYDVR0RBBQwEoIQYmxvWHJvdXRlLmRldi5DQTAMBgNVHRMEBTADAQH/\nMAsGA1UdDwQEAwIB/jAKBggqhkjOPQQDAgNpADBmAjEAo6kPPChOytP961lFjKFb\n+zfEPm6sHtBxmgeDMhQwqb1erIIsYfU6zVaA82g9REHvAjEAoLfzcjEq91/Jlcmn\nCSgJY3JUPIocBek+o9cKczwz1ZDuzGscMOF0J4fpTwAyJOUP\n-----END CERTIFICATE-----"
