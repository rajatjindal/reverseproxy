# reverseproxy

The usecase for this proxy is:  

- when you have a service (service-a) running on your local setup with self-signed certificate and
- you need this service (service-a) from other service (service-b), but that other service (service-b) does not care if your service (service-a) is http or https but
- the service-b needs a valid certificate if the service-a is https url
