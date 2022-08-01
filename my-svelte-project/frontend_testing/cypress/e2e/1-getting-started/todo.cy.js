/// <reference types = "cypress" />


describe('todo frontend testing', () => {
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

    it('pressing on the all button and check todos', () => {
        cy.get('#btn2').click()
        cy.contains('cypress edit test').should('exist')
        cy.contains('cypress testing 2').should('exist')
    })  

    it('pressing on the active button and check todos', () => {
        cy.get('#btn3').click()
        cy.contains('cypress edit test').should('not.exist')
        cy.contains('cypress testing 2').should('exist')
    }) 
    it('pressing on the completed button and check todos', () => {
        cy.get('#btn4').click()
        cy.contains('cypress edit test').should('exist')
        cy.contains('cypress testing 2').should('not.exist')
    }) 
    it('delete all checked todos button' , () => {      
        cy.get('#btn10').click()
    })
    it('check first todo is not found , while second is found' , () => {      
        cy.contains('cypress edit test').should('not.exist')
        cy.contains('cypress testing 2').should('exist')
    })
    it('testing adding new todo' , () => {
        cy.get('#todo-0').type("cypress testing 3")
        cy.get('#btn1').click()
        cy.contains("cypress testing 3")
    })
    
    it('testing check all button' , () => {
        cy.get('#btn9').click()
    })


    it('pressing on the all button and check todos', () => {
        cy.get('#btn2').click()
        cy.contains('cypress testing 2').should('exist')
        cy.contains('cypress testing 3').should('exist')
    })  

    it('pressing on the active button and check todos', () => {
        cy.get('#btn3').click()
        cy.contains('cypress testing 2').should('not.exist')
        cy.contains('cypress testing 3').should('not.exist')
    }) 
    it('pressing on the completed button and check todos', () => {
        cy.get('#btn4').click()
        cy.contains('cypress testing 2').should('exist')
        cy.contains('cypress testing 3').should('exist')
    }) 

    it('delete all todos button' , () => {      
        cy.get('#btn10').click()
    })

    it('checking there is no todos left', () => {
        cy.contains('cypress testing 2').should('not.exist')
        cy.contains('cypress testing 3').should('not.exist')
    }) 

  

})