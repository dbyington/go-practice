package queue

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", func() {
    var err error
    var testQ queue
    var qr, testQr Queuer
    var qSize uint64
    var testItem queueItem

    BeforeEach(func() {
        qSize = 5
        qr = NewQueuer(qSize)
        testQ = queue{
            list: []queueItem{},
            size: qSize,
        }
        testQr = Queuer(&testQ)
    })

    Describe(".NewQueuer", func() {
        It("should be a Queuer", func() {
            Expect(qr).To(BeEquivalentTo(testQr))
        })
    })

    Describe("#Enqueue", func() {
        Context("with the first item", func() {
            BeforeEach(func() {
                testItem = 42
                err = testQr.Enqueue(testItem)
            })

            It("should not error", func() {
                Expect(err).ToNot((HaveOccurred()))
            })

            It("should add the item to the queue", func() {
                Expect(testQ.list[0]).To(Equal(testItem))
            })
        })

        Context("when the queue is full", func() {
            BeforeEach(func() {
                testQ.list = []queueItem{42, 37, 28, 17, 51}
                err = testQr.Enqueue(testItem)
            })

            It("should return an error", func() {
                Expect(err).To(MatchError(ErrQueueFull))
            })
        })
    })

    Describe("#Dequeue", func() {
        var dqItem queueItem
        Context("with an item to dequeue", func() {
            BeforeEach(func() {
                testItem = 42
                testQ.list = []queueItem{testItem, 37}
                dqItem, err = testQr.Dequeue()
            })

            It("should not error", func() {
                Expect(err).ToNot(HaveOccurred())
            })

            It("should dequeue return the bottom item from the queue", func() {
                Expect(dqItem).To(BeEquivalentTo(testItem))
            })
        })

        Context("with an empty queue", func() {
            BeforeEach(func() {
                testQ.list = []queueItem{}
                dqItem, err = testQr.Dequeue()
            })

            It("should error", func() {
                Expect(err).To(MatchError(ErrQueueEmpty))
            })
        })
    })

    Describe("#Peek", func() {
        var pItem queueItem
        Context("with an item in the queue", func() {
            BeforeEach(func() {
                testItem = 42
                testQ.list = []queueItem{testItem}
                pItem, err = testQr.Peek()
            })

            It("should not error", func() {
                Expect(err).ToNot(HaveOccurred())
            })

            It("should return the item", func() {
                Expect(pItem).To(BeEquivalentTo(testItem))
            })
        })

        Context("with an empty queue", func() {
            BeforeEach(func() {
                testQ.list = []queueItem{}
                pItem, err = testQr.Peek()
            })

            It("should error", func() {
                Expect(err).To(MatchError(ErrQueueEmpty))
            })
        })
    })
})
