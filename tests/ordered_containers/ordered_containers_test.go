package orderedcontainers

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Ordered", Ordered, func() {
	// Runs before all the child specs (It) in the container (It1, It2, It3)
	BeforeAll(func() {
		fmt.Println("BeforeAll Describe")
	})

	// Runs after all the child specs (It) in the container (It1, It2, It3)
	AfterAll(func() {
		fmt.Println("AfterAll Describe")
	})

	// Runs before each spec (It) in the Describe container (It1, It2, It3)
	BeforeEach(func() {
		fmt.Println("BeforeEach Describe")
	})

	// Runs after each spec (It) in the Describe container (It1, It2, It3)
	AfterEach(func() {
		fmt.Println("AfterEach Describe")
	})

	Context("Context1", func() {
		// Runs before all the child specs (It) in the container (It1)
		BeforeAll(func() {
			fmt.Println("BeforeAll Context1")
		})

		// Runs after all the child specs (It) in the container (It1)
		AfterAll(func() {
			fmt.Println("AfterAll Context1")
		})

		// Runs before each spec (It) in the Context container (It1)
		BeforeEach(func() {
			fmt.Println("BeforeEach Context1")
		})
	
		// Runs after each spec (It) in the Context container (It1)
		AfterEach(func() {
			fmt.Println("AfterEach Context1")
		})

		It("It1", func() {
			fmt.Println("It1")
		})
	})

	Context("Context2", func() {
		// Runs before all the child specs (It) in the container (It2 and It3)
		BeforeAll(func() {
			fmt.Println("BeforeAll Context2")
		})

		// Runs after all the child specs (It) in the container (It2 and It3)
		AfterAll(func() {
			fmt.Println("AfterAll Context2")
		})

		// Runs before each spec (It) in the Context container (It2 and It3)
		BeforeEach(func() {
			fmt.Println("BeforeEach Context2")
		})
	
		// Runs after each spec (It) in the Context container (It2 and It3)
		AfterEach(func() {
			fmt.Println("AfterEach Context2")
		})

		When("When1", func() {
			// Runs before each spec (It) in the When container (It2)
			BeforeEach(func() {
				fmt.Println("BeforeEach When1")
			})
		
			// Runs after each spec (It) in the When container (It2)
			AfterEach(func() {
				fmt.Println("AfterEach When1")
			})

			It("It2", func() {
				fmt.Println("It2")
			})
		})

		It("It3", func() {
			fmt.Println("It3")
		})
	})
})
