# mrcachego
In memory KV cache implemented in go

As this is 5x slower than mrcache written in C I haven't bothered to finish this up. It is currently using gnet which created a fake benchmark to make its performance look far better than it is.  I'll come back to this if go gets a decently performing tcp stack. 
