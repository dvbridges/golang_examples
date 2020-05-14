# Gonum real-time plots

The main.go code was taken from [justforfunc](https://www.youtube.com/watch?v=77ZFKuMLkp4).

This code demonstrates how to create dynamic graphs using Gonum, HTML and JS.\
We measure the endpoint latency time to load localhost:8080, and plot the latencies using a scatter\
plot and their distribution using a histogram. To automate the visits to the localhost, use the\
following from Linux terminal
```while true; do curl localhost:8080; done
```





