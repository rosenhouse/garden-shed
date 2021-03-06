package repository_fetcher_test

import (
	"errors"
	"net/url"

	"github.com/cloudfoundry-incubator/garden-shed/layercake"
	"github.com/cloudfoundry-incubator/garden-shed/repository_fetcher"
	fakes "github.com/cloudfoundry-incubator/garden-shed/repository_fetcher/repository_fetcherfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/lager/lagertest"
)

var _ = Describe("Retryable", func() {
	var (
		fakeRemoteFetcher *fakes.FakeRepositoryFetcher

		repoURL   *url.URL
		logger    *lagertest.TestLogger
		retryable repository_fetcher.Retryable
	)

	BeforeEach(func() {
		var err error
		fakeRemoteFetcher = new(fakes.FakeRepositoryFetcher)

		repoURL, err = url.Parse("http://fake-registry-1.docker.io/")
		Expect(err).NotTo(HaveOccurred())
		logger = lagertest.NewTestLogger("test")

		retryable = repository_fetcher.Retryable{
			RepositoryFetcher: fakeRemoteFetcher,
			Logger:            logger,
		}
	})

	Describe("Fetch failures", func() {
		Context("when fetching fails twice", func() {
			BeforeEach(func() {
				fakeRemoteFetcher.FetchStub = func(u *url.URL, diskQuota int64) (*repository_fetcher.Image, error) {
					if fakeRemoteFetcher.FetchCallCount() <= 2 {
						return nil, errors.New("error-talking-to-remote-repo")
					} else {
						return nil, nil
					}
				}

				_, err := retryable.Fetch(repoURL, 0)
				Expect(err).NotTo(HaveOccurred())
			})

			It("suceeds on third attempt", func() {
				Expect(fakeRemoteFetcher.FetchCallCount()).To(Equal(3))
			})

			It("logs failing attempts", func() {
				itLogsFailingAttempts(logger, 2, "test.failed-to-fetch")
			})
		})

		Context("when fetching fails three times", func() {
			BeforeEach(func() {
				fakeRemoteFetcher.FetchStub = func(u *url.URL, diskQuota int64) (*repository_fetcher.Image, error) {
					return nil, errors.New("error-talking-to-remote-repo")
				}
				_, err := retryable.Fetch(repoURL, 0)
				Expect(err).To(HaveOccurred())
			})

			It("returns an error", func() {
				Expect(fakeRemoteFetcher.FetchCallCount()).To(Equal(3))
			})

			It("logs failing attempts", func() {
				itLogsFailingAttempts(logger, 3, "test.failed-to-fetch")
			})
		})
	})

	Describe("FetchID failures", func() {
		Context("when fetching IDs fails twice", func() {
			BeforeEach(func() {
				fakeRemoteFetcher.FetchIDStub = func(u *url.URL) (layercake.ID, error) {
					if fakeRemoteFetcher.FetchIDCallCount() <= 2 {
						return nil, errors.New("error-talking-to-remote-repo")
					} else {
						return nil, nil
					}
				}

				_, err := retryable.FetchID(repoURL)
				Expect(err).NotTo(HaveOccurred())
			})

			It("suceeds on third attempt", func() {
				Expect(fakeRemoteFetcher.FetchIDCallCount()).To(Equal(3))
			})

			It("logs failing attempts", func() {
				itLogsFailingAttempts(logger, 2, "test.failed-to-fetch-ID")
			})
		})

		Context("when fetching IDs fails three times", func() {
			BeforeEach(func() {
				fakeRemoteFetcher.FetchIDStub = func(u *url.URL) (layercake.ID, error) {
					return nil, errors.New("error-talking-to-remote-repo")
				}

				_, err := retryable.FetchID(repoURL)
				Expect(err).To(HaveOccurred())
			})

			It("returns an error", func() {
				Expect(fakeRemoteFetcher.FetchIDCallCount()).To(Equal(3))
			})

			It("logs failing attempts", func() {
				itLogsFailingAttempts(logger, 3, "test.failed-to-fetch-ID")
			})
		})
	})
})

var itLogsFailingAttempts = func(logger *lagertest.TestLogger, count int, msg string) {
	Expect(len(logger.LogMessages())).To(Equal(count))
	for _, logMsg := range logger.LogMessages() {
		Expect(logMsg).To(Equal(msg))
	}
}
