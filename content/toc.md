---
title: "Table of Contents"
date: 2021-12-08T08:27:41+11:00
---

A PDF of the detailed table of contents is available for download from the book's
product page [here](https://media.wiley.com/product_data/excerpt/14/11197738/1119773814-6.pdf).

Here's a brief description of what you can expect to find in each chapter.

The book consits of eleven main chapters, another three other auxiliary chapters - Getting started and appendices. 

In **Getting Started**, I guide you through the installation of the software required for the
book. I also discuss some of the assumed knowledge for the following chapters.

In **Chapters 1 and 2**, I show you how to write **command line applications** using the
standard library's `flag` package. You will learn to write testable comand line applications,
implement sub-commands in your application as well as handling timeouts and user signals.
You can read an excerpt from Chapter 1 [here](https://media.wiley.com/product_data/excerpt/14/11197738/1119773814-39.pdf).

In **Chapters 3 and 4**, I show you how you can write robust **HTTP clients**, using the
standard library's `net/http` package and friends. You will learn how to marshal and unmarshal
data, enforcing timeouts, writing client middleware and connection pooling.

In **Chapters 5, 6 and 7**, I show you how you can build production ready **HTTP servers**
using `net/http` and related packages from the standard library. You will learn to handle streaming data, write middleware, sharding data across handler functions, 
handle client disconnections and learn strategies to implement timeouts at
various points in your server application.

In **Chapters 8, 9 and 10**, I show you can build production ready **gRPC applications**
using the Go gRPC implementation. You will learn to implement unary and streaming RPC communication patterns, implement middleware using intereceptors, improving the 
resiliency of your applications and securing communication using TLS certificates.

In **Chapter 11**, I show you can work with two categories of **data stores** - Object store
and a SQL data store using the https://gocloud.dev library. You will learn to implement
vendor agnostic applications and also learn to test your applications.

In **Apendix A**, you will learn to implement instrumentation in applications to emit
telemetry data - Logs, Metrics, and Traces. You will learn to integrate OpenTelemetry in
your applications in this appendix.

In **Appendix B**, you will learn to read configuration data in your application, learn
to build a container image for your application and get some guidance around deploying
your applications.