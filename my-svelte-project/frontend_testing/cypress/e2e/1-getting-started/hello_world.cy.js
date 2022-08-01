/// <reference types = "cypress" />

describe('basic test', () => {
    it('testing valid header' , () => {
        cy.visit('http://localhost:3000')
        cy.contains("Svelte to-do list")
        cy.get()
    })
})