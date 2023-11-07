# About exosip2

This provides binding to the GNU osip2 library and the eXosip2 stack. This will
allow writing of event driven SIP applications in the Go language.  The eXosip2
stack has a high level of interop conformance and has been extensively used for
over two decades in both open source projects and commercial products.

This implimentation offers a very small subset of eXosip2 functionality both as
proof of concept for integrating eXosip2 with Go, and to jump-start further
development.  There is a lot of cgo bridging code, especially for event
structures.  However, given that other open source SIP conformat stacks in pure
Go seem either to still be very incomplete or have been entirely abandoned many
years go, this may actually prove more viable for Go SIP development.

