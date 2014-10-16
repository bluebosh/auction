package auction_http_handlers_test

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/cloudfoundry-incubator/auction/auctiontypes"
	"github.com/cloudfoundry-incubator/auction/communication/http/routes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RebidThenReserveHandler", func() {
	Context("with valid JSON", func() {
		var startAuctionInfo auctiontypes.StartAuctionInfo

		BeforeEach(func() {
			startAuctionInfo = auctiontypes.StartAuctionInfo{
				ProcessGuid:  "process-guid",
				InstanceGuid: "instance-guid",
				DiskMB:       1024,
				MemoryMB:     256,
				Index:        1,
			}
		})

		It("should notify the auction rep", func() {
			Request(routes.RebidThenTentativelyReserve, nil, JSONReaderFor(startAuctionInfo))
			Ω(auctionRep.RebidThenTentativelyReserveCallCount()).Should(Equal(1))
			Ω(auctionRep.RebidThenTentativelyReserveArgsForCall(0)).Should(Equal(startAuctionInfo))
		})

		Context("and a succesful bid and reservation", func() {
			BeforeEach(func() {
				auctionRep.RebidThenTentativelyReserveReturns(0.73, nil)
			})

			It("should return the response of the fake", func() {
				status, body := Request(routes.RebidThenTentativelyReserve, nil, JSONReaderFor(startAuctionInfo))
				Ω(status).Should(Equal(http.StatusOK))
				Ω(body).Should(MatchJSON(JSONFor(auctiontypes.StartAuctionBid{
					Rep:   repGuid,
					Bid:   0.73,
					Error: "",
				})))
			})
		})

		Context("and an unsuccesful bid and reservation", func() {
			BeforeEach(func() {
				auctionRep.RebidThenTentativelyReserveReturns(0, errors.New("oops"))
			})

			It("should return a non-happy status code and the error", func() {
				status, body := Request(routes.RebidThenTentativelyReserve, nil, JSONReaderFor(startAuctionInfo))
				Ω(status).Should(Equal(http.StatusForbidden))
				Ω(body).Should(MatchJSON(JSONFor(auctiontypes.StartAuctionBid{
					Rep:   repGuid,
					Bid:   0,
					Error: "oops",
				})))
			})
		})
	})

	Context("when invalid JSON", func() {
		It("should return an error without calling the rep", func() {
			status, body := Request(routes.RebidThenTentativelyReserve, nil, bytes.NewBufferString("∆"))
			Ω(status).Should(Equal(http.StatusBadRequest))
			Ω(body).Should(ContainSubstring("invalid json: invalid character"))

			Ω(auctionRep.RebidThenTentativelyReserveCallCount()).Should(BeZero())
		})
	})
})