package main

func history() (chan<- PackedMessage, func(string) []map[string]interface{}) {
	buffer := make(chan PackedMessage, 100)
	channels := map[string][]map[string]interface{}{}

	go func() {
		for h := range buffer {
			if _, exists := channels[h.Channel]; !exists {
				channels[h.Channel] = []map[string]interface{}{}
			}

			// Shift first item in the historic
			if len(channels[h.Channel]) >= 100 {
				channels[h.Channel] = channels[h.Channel][1:]
			}

			channels[h.Channel] = append(channels[h.Channel], h.Message)
		}
	}()

	last := func(channel string) []map[string]interface{} {
		if list, exists := channels[channel]; exists {
			return list
		}

		return []map[string]interface{}{}
	}

	return buffer, last
}
