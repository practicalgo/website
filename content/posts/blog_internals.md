---
title:  How is practicalgobook.net reaching you?
date: 2022-07-27
categories:
-  articles
draft: true
---
I write my blog posts in Markdown format,  and then use the excellent [Hugo](https://gohugo.io/) static site 
generator to generate HTML files. The topic of this blog post is how those HTML files are served
to you.

When you type in https://practicalgobook.net in your browser (or click it from somewhere), a translation,
commonly known as DNS resolution happens. As a result of this translation, the browser gets back
an IP address. This IP address to a virtual machine running in the cloud. When the browser's HTTPS request
reaches this virtual machine, a program then sends back the content that you read on this website. In this
post, I am going to describe that program. It is written using the Go programming language using standard
libraries, with some reimplementation of logic from third party packages. Let's dive in!

## Content delivery


## TLS certificate management
