package main

const (
	// Key server
	Key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAtH8//pyL/Ms82qdZiAZaOnCgnFtdcirifDEoVXpbA881gezS
lo2NDfM9FO05BUiAeAwoDYsxVrPudr00d9wGaZaNgxyyofiusT8Amo7qmrk4RGFa
JDK2hsHfacSY/2mu5ktWX43hf4kxLkvfOKJ8bqs40AyU+Enyvavz5av79H3xTQtT
YXzDSVRtEYn2F9zK4f7AIAKveHXXBI06ibepNpBTuQK1zksDtXP1jFVyFY2OLJEn
mJwYRfwgP6dLKhMpFuyQfWq8Tl5tVT10/StNMsupVfeT0C1yDejTWPYG0Gi9cmwV
7MmLJfgOQcb+FghLltFZSGkxsECr+PbGDafywQIDAQABAoIBAQCtikdaIuPJKxMY
W5yOVyU7WPxjOMkQssDNjWntPVXyxK/6eltDG1fKWNqamzaeqMW0JmQAGZtXdGpL
ScD/mNGlhXYL4HSznGlBy4RmFhDksD60KDHF6ehqik6m37oGoHLjwv2Jo91yhNJ/
Y718M2WV9hQKnH5Cdr7ntEMsyH3X4/xmHyBFuVP99siyNCjRNCUxUWtvNCL1RqXP
T1xyo3kC1Bfh5F5EnO3ESbF1Hsn1X/MKDeW/DJDxinh3jbgx9zd0MQWb1YBvVspn
YaHUpWpkb+XJ0jdxeF232r6j2VNFk8Lz7mcc1mDhunLDNrCt5OHD5L67oM4HG7Rj
QxiS/b0NAoGBAOm9fAONl4FSOcDGmWDQoqYLzK4XmkUGaBZcKXZ3fYVvR3jNZWcO
GbbxYaoz9h7bVu3AktgftDhXYHfQB89i+GnRP6Ry0+9b1EzLIXBBGXAXiG6gXTMU
3xoSIeEVx6pZXoUyVRFBDIgqcoVPGv4+AqQGRubIwsAVOitOqmJwAq7fAoGBAMWv
tvBIq27pHvoVN9s++cvCIIQjO2t5zZSgIR752L+HL4ZerwRCfDdOLxN7kkF+IJJ6
dhLu189+XSXf1QxaaDkED0IpJyZwYRgTdbUm17/fDhhokIU9uN1x7M84+saKZAy9
O/9dJ2vPW5rj+6q7+n3M1Jcq9IkxxqqIR1wtLrJfAoGBANgXy0hJphD0INqdgP1l
xYk3jXJB9ejspFxPtjGFBPHQ0EXZtm7RNWrRvFYrldYvOC20BM50eRxBg7khBadw
u3Kw9mKlmlRHFH4uqepq9QaRaxvDfIaWPDCRJOtARIiz+NOxlb3O4rQcciXW3YDB
eZWFDBWe85W7yjxxjMpW/dQHAoGAav31FwtdrYOCfnupZYIqFDuW+a1P5ZVgzMX1
5xv2UYLLGAB63OVW8V/hXrwMpmGrI1wgN5MJPoX3yB3i4vKzYdhuobJAPC+qfStb
E5ZYQsJokJFXVqXXPhMAxg7iuz3/dF2e8VDrmw/5VlV5yig/JIMNtFtNSkukaHul
k5/onCUCgYEAp/b55AH1C1WzMCtmQQELNDKFXbhlnUPimw7HNafAxIjJ9IueMHjR
yqof4rLl/eqnNTZI4gXxt2ZbR1QkQs2IUzRrF59PMR85FaiQhIDY9rELP2fwQCHr
Fq3Y5Kk7BIenWqNmMzoV1V3dDgjnYxP/n+Jc1aebSu0ujYh+OKXxcDc=
-----END RSA PRIVATE KEY-----`
	// Cert server
	Cert = `-----BEGIN CERTIFICATE-----
MIID8TCCAtmgAwIBAgIJAIiVuyZvrWWjMA0GCSqGSIb3DQEBCwUAMIGOMQswCQYD
VQQGEwJCUjEPMA0GA1UECAwGQW1hem9uMQ8wDQYDVQQHDAZNYW5hdXMxFzAVBgNV
BAoMDlJvZG9sZm8ncyBIb21lMQ0wCwYDVQQLDARIb21lMRIwEAYDVQQDDAlsb2Nh
bGhvc3QxITAfBgkqhkiG9w0BCQEWEnJvZjIwMDA0QGdtYWlsLmNvbTAeFw0xNzAz
MDExODA5MDBaFw0yNzAyMjcxODA5MDBaMIGOMQswCQYDVQQGEwJCUjEPMA0GA1UE
CAwGQW1hem9uMQ8wDQYDVQQHDAZNYW5hdXMxFzAVBgNVBAoMDlJvZG9sZm8ncyBI
b21lMQ0wCwYDVQQLDARIb21lMRIwEAYDVQQDDAlsb2NhbGhvc3QxITAfBgkqhkiG
9w0BCQEWEnJvZjIwMDA0QGdtYWlsLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBALR/P/6ci/zLPNqnWYgGWjpwoJxbXXIq4nwxKFV6WwPPNYHs0paN
jQ3zPRTtOQVIgHgMKA2LMVaz7na9NHfcBmmWjYMcsqH4rrE/AJqO6pq5OERhWiQy
tobB32nEmP9pruZLVl+N4X+JMS5L3ziifG6rONAMlPhJ8r2r8+Wr+/R98U0LU2F8
w0lUbRGJ9hfcyuH+wCACr3h11wSNOom3qTaQU7kCtc5LA7Vz9YxVchWNjiyRJ5ic
GEX8ID+nSyoTKRbskH1qvE5ebVU9dP0rTTLLqVX3k9Atcg3o01j2BtBovXJsFezJ
iyX4DkHG/hYIS5bRWUhpMbBAq/j2xg2n8sECAwEAAaNQME4wHQYDVR0OBBYEFDku
bWsvckdqjMhA6azn+/iBZe8kMB8GA1UdIwQYMBaAFDkubWsvckdqjMhA6azn+/iB
Ze8kMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAFwVf2rO3uRiYQRn
SYNrlQw0l3f+5XuDS9lL6QMxoV1pt1RmiYJDxk1Bz78Bc7bqDlrvohre7r3xsbOX
qzIDnCu2ZBm0QtjLwihNbTQMxj/2RxIPhSACB/3UlFLRlupHUAwMFsM3h2oJ2Im4
HMagUEqQBW5it9FaAQgL3cUYfcISLHx9dIQ57VHAtD2kKCA3XKAy4U7eMusKeg/e
TGihJ1y4df4X5rgX89IjFfn8V27SdioHUoej/JazBsvTefmVMTeR3kUXyS8Zs8af
u9HLJbdqPFoty3AzMoYfnOnF6QN9gxoHB34RWyFgEbLo437cdMsgZIw0B2+Q74fC
A9e0BlA=
-----END CERTIFICATE-----`
)
