/// <reference types = "cypress" />

describe('basic test', () => {
    it('testing valid header' , () => {
        cy.visit('http://localhost:3000')
        cy.contains("Svelte to-do list")
      
    })

    it('testing adding two new todo' , () => {
        cy.visit('http://localhost:3000')
        cy.get('#todo-0').type("cypress testing 1")
        cy.get('#btn1').click()
        cy.get('#todo-0').type("cypress testing 2 ")
        cy.get('#btn1').click()
        cy.contains("cypress testing")
    })

    it('testing check all button' , () => {
        cy.visit('http://localhost:3000')
        cy.get('#btn9').click()
    })

    it('delete all todos button' , () => {
        cy.visit('http://localhost:3000')
        cy.get('#btn10').click()
    })

    it('testing adding two new todo' , () => {
        cy.visit('http://localhost:3000')
        cy.get('#todo-0').type("cypress testing 1")
        cy.get('#btn1').click()
        cy.get('#todo-0').type("cypress testing 2 ")
        cy.get('#btn1').click()
        cy.contains("cypress testing")
    })

    it('testing pressing on all and active and completed' , () => {
        cy.visit('http://localhost:3000')
        cy.get('#btn2').click()
        cy.get('#btn3').click()
        cy.get('#btn4').click()
    })

})