# All about channel
Describe source code
1. Sample from [Channel Axioms](https://dave.cheney.net/2014/03/19/channel-axioms)
   * When you send or received from nil channel -> deadlock
   * When received from closed channel -> return 0 when channel drain all of this element