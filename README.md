# kubeauth


`kubeauth` is a OIDC (OpenId Connect) helper forked from [dexter](https://github.com/gini/dexter) to create a hassle-free Kubernetes login experience.
This fork adds okta fat id_token support, enabling RBAC rolebindings based on groups.

This fork disables the original google and azure authentication; for those implementations, please see the original `dexter` repo.

## Authentication Flow

`kubeauth` will open a new browser window and redirect you to Okta. The only interaction you have is the Okta login and your kubectl k8s config is updated automatically.

## Okta and RBAC Setup
The Okta/RBAC integration process involves:   
- setting up an Okta Authorization Server;   
- setting up an Okta OIDC app;   
- configuring the k8s cluster kubeapiserver.   

(To be added... Detailed description for setting up Okta OIDC and Authentication servers and the corresponding RBAC kubeapi configuration).   

[kubeapi server config flags](https://zoom.us/j/7533652875?pwd=RlBSOVZNTjRKdTArMUNqVmc5SkRmZz09)    
[okta authorization server](https://developer.okta.com/docs/concepts/auth-servers/).  
[okta oidc app](https://developer.okta.com/docs/concepts/auth-overview/#openid-connect).  

## Run kubeauth

Run `kubeauth` without a command to access the help screen/intro.

```
❯ ./build/kubeauth_darwin_amd64
 _          _                      _   _     
| |        | |                    | | | |    
| | ___   _| |__   ___  __ _ _   _| |_| |__  
| |/ / | | | '_ \ / _ \/ _` | | | | __| '_ \ 
|   <| |_| | |_) |  __/ (_| | |_| | |_| | | |
|_|\_\\__,_|_.__/ \___|\__,_|\__,_|\__|_| |_|
kubeauth is a authentication helper for Kubernetes that does the heavy
lifting for SSO (Single Sign On) for Kubernetes.

Usage:
  kubeauth [command]

Available Commands:
  auth        Authenticate with OIDC provider
  help        Help about any command
  version     Print the version number of kubeauth

Flags:
  -h, --help      help for kubeauth
  -v, --verbose   verbose output

Use "kubeauth [command] --help" for more information about a command.
```

Running `kubeauth auth` will start the authentication process.

```
 ❯ ./build/kubeauth_darwin_amd64 auth --help
Use your Okta login to get a JWT (JSON Web Token) and update your
local k8s config accordingly. A refresh token is added and automatically refreshed
by kubectl. Existing token configurations are overwritten.

For details on the original implementation go to: https://blog.gini.net/

kubeauth authentication flow
===========================

1. Open a browser window/tab and redirect you to Google (https://accounts.google.com)
2. You login with your Google credentials
3. You will be redirected to kubeauths builtin webserver and can now close the browser tab
4. kubeauth extracts the token from the callback and patches your ~/.kube/config

Usage:
  kubeauth auth [flags]

Flags:
  -c, --callback string        Callback URL. The listen address is dreived from that. (default "http://127.0.0.1:64464/callback")
  -i, --client-id string       Google clientID (default "REDACTED")
  -s, --client-secret string   Google clientSecret (default "REDACTED")
  -d, --dry-run                Toggle config overwrite
  -e, --endpoint string        OIDC-providers: google or azure (default "google")
  -h, --help                   help for auth
  -k, --kube-config string     Overwrite the default location of kube config (~/.kube/config) (default "/Users/dkerwin/.kube/config")
  -t, --tenant string          Your azure tenant (default "common")

Global Flags:
  -v, --verbose   verbose output
```


# From the original README: ===>

## Configuration
### Google credentials

  -  Open [console.developers.google.com](https://console.developers.google.com)
  -  Create new credentials
      - OAuth Client ID
      - Web Application
      - Authorized redirect URIs: http://127.0.0.1:64464/callback

### Or, configure Azure credentials

  -  Open [portal.azure.com](https://portal.azure.com)
  -  Go to App registrations and create a new app
      - Enter reply URI http://127.0.0.1:64464/callback
      - Create secret key
      - Collect application ID (client ID)

### Auto pilot configuration
kubeauth also support auto pilot mode. If your existing kubectl context uses one of the supported OIDC-providers, kubeauth will try to use the OIDC details from kubeconfig.

## Installation



It is possible to embed your Google credentials into the resulting binary.

```
CLIENT_ID=abc123.apps.googleusercontent.com CLIENT_SECRET=mySecret OS=linux make
```

## Contribution Guidelines

It's awesome that you consider contributing to `kubeauth` and it's really simple. Here's how it's done:

  - fork repository on Github
  - create a topic/feature branch
  - push your changes
  - update documentation if necessary
  - open a pull request

## Authors

Initial code was written by [Daniel Kerwin](mailto:daniel@gini.net) & [David González Ruiz](mailto:david@gini.net)   
`kubeauth` fork was written by [David Rosenbloom](mailto:david.rosenbloom@pm.me)

## Acknowledgements

`dexter` was inspired by this [blog post series](https://thenewstack.io/tag/Kubernetes-SSO-series) by [Joel Speed](https://thenewstack.io/author/joel-speed/), [Micah Hausler's k8s-oidc-helper
](https://github.com/micahhausler/k8s-oidc-helper) & [CoreOS dex](https://github.com/coreos/dex).

## License

MIT License. See [License](/LICENSE) for full text.

