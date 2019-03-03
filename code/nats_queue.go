nc.QueueSubscribe("help", "job_workers", func(m *Msg) { // HL
	nc.Publish(m.Reply, []byte("I can help!"))
})
