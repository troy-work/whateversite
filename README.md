# whateversite
golang metro bootstrap templates https responsive

Please don't start a server like I did. I'm a noob with go. The ListenAndServe. . .can't be shut down. I'll fix this shortly. . . 

This redirects http to https, so you'll have to permit the included unverified self-signed certs to run locally.
Obviously you need to replace these with your own verified certs to use it in production. 

It links in cdn resources for all the metro and jquery stuff in the header.tmpl (template).
You can easily change these if you need to support ie8. As it is, it fully supports a responsive design.

It's called whatever because you will never go to a page that doesn't exist, meaning every page exists when you go to it. And every page is editable.

It's an example of view handlers that do not require pr-defined routes.

