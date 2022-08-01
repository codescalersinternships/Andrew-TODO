/// <reference types = "cypress" />


describe('basic test', () => {
    beforeEach(() => {
        cy.visit('http://localhost:3000')
      })
    it('testing valid header' , () => {    
        cy.contains("Svelte to-do list")
    })

    it('testing adding two new todo' , () => {
        cy.get('#todo-0').type("cypress testing 1")
        cy.get('#btn1').click()
        cy.get('#todo-0').type("cypress testing 2")
        cy.get('#btn1').click()
        cy.contains("cypress testing 1")
        cy.contains("cypress testing 2")
    })

    it('checking the two todos', () => {
        cy.get('#todo-text-1').should('have.text', 'cypress testing 1')
        cy.get('#todo-text-2').should('have.text', 'cypress testing 2')
      })

    it('edit the first todo', () => {
        cy.get('#btn7-1').click()
        cy.get('#todo-edit-1').type("cypress edit test")
        cy.get('#btn6-1').click()
      
      })
    it('checking the two todos', () => {
        cy.get('#todo-text-1').should('have.text', 'cypress edit test')
        cy.get('#todo-text-2').should('have.text', 'cypress testing 2')
    }) 
    
    it('edit the second  todo "press cancel not save button"', () => {
        cy.get('#btn7-1').click()
        cy.get('#todo-edit-1').type("cypress edit test")
        cy.get('#btn5-1').click()
      })

    it('checking the two todos', () => {
        cy.get('#todo-text-1').should('have.text', 'cypress edit test')
        cy.get('#todo-text-2').should('have.text', 'cypress testing 2')
    })  

    it('checking the first todo as done ', () => {
        cy.get('#todo-check-1').click()
    })  
    
    it('delete all checked todos button' , () => {      
        cy.get('#btn10').click()
    })

    it('testing pressing on all and active and completed' , () => {
          cy.get('#btn2').click()
          cy.get('#btn3').click()
          cy.get('#btn4').click()
      })
        
    it('testing check all button' , () => {
        cy.get('#btn9').click()
    })

    it('delete all todos button' , () => {      
        cy.get('#btn10').click()
    })


  

})