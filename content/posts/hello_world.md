---
title:  Hello world
date: 2021-12-22
categories:
-  updates
---

Hello everyone. "Practical Go: Building Scalable Network and Non-Network Applications" is now out 
electronically as well as in print. 

The book targets readers who have learned the Go language basics - syntax, understanding of types, 
being able to write functions (details are available in the "Getting started" chapter - but essentially,
just a passing familiarity with the language is enough I think). I say that and reasonably confidently
since that's where I was when I started writing this book. I wrote this book, not to share with you things 
I already knew - far from it. The book is essentially a documentation of my journey of learning Go and applying
it to solve problems which I think based on my experience are being solved by folks in the industry a.k.a
real world today.

The book consists of 11 chapters, a Getting Started guide and two Appendices.


In Getting Started, I guide you through the installation of the software required for the book. I also discuss some of the assumed knowledge for the following chapters.

In Chapters 1 and 2, I show how to write command line applications using the standard library’s flag package.

In Chapters 3 and 4, I show how you can write robust HTTP clients, using the standard library’s net/http package and friends.

In Chapters 5, 6 and 7, I show how you can build production ready HTTP servers using net/http
and related packages from the standard library.

In Chapters 8, 9 and 10, I show how you can build production ready gRPC applications using the Go gRPC implementation.

In Chapter 11, I show how you can work with two categories of data stores - Object store and a SQL data store using the https://gocloud.dev library.

In Apendix A, you will learn to implement instrumentation in applications to emit telemetry data - Logs, Metrics, and Traces.

In Appendix B, you will learn to read configuration data in your application, learn to build a container image for your application and get some guidance around deploying your applications.

You can find a detailed table of contents here: https://practicalgobook.net/toc/ along with a link to an excerpt of Chapter 1. 
Any further questions, let me know, i will be happy to answer!

You can subscribe to the blog using your favorite [RSS reader](https://practicalgobook.net/index.xml). You can also follow
the book's [Twitter account](https://twitter.com/practicalgobook).

In the next post, I will share how this website is hosted! Hint: it's a Go HTTP Server.
