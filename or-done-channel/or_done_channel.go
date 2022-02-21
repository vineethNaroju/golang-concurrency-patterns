package ordone

func OrDone(doneChan, inputStream <-chan interface{}) <-chan interface{} {
	results := make(chan interface{})

	go func() {
		defer close(results)

		for {
			select {
			case <-doneChan:
				return
			default:
				val, ok := <-inputStream

				if !ok {
					return
				}

				select {
				case <-doneChan:
					return
				default:
					results <- val
				}
			}
		}
	}()

	return results
}
