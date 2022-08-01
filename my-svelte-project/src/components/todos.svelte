
<script>
  import axios from 'axios'
  import {onMount} from 'svelte'
    $:todos = []
    onMount(async () => {
      try {
	        const res = await axios.get('http://localhost:8080/todos');
          todos = await res.data
	    } catch (e) {
		        console.log(e)
	    }
});
const remove_todo_api = async (id) => {
		try {
      const res = await axios.delete('http://localhost:8080/todos/'+id);
			console.log("should be deleted")
		} catch (err) {
			console.log(err);
		}
	};

 

const add_todo_api = async (item) => {
		try {
      item = JSON.stringify(item)
      const res = await axios.post('http://localhost:8080/todos', item);
			console.log("should be deleted")
		} catch (err) {
			console.log(err);
		}
	};  
  const update_todo_api = async (id,item) => {
		try {
      item = JSON.stringify(item)
      const res = await axios.put('http://localhost:8080/todos/'+id, item);
		} catch (err) {
			console.log(err);
		}
	};   

  const update_todo_status_api = async (id) => {
		try {
      const res = await axios.patch('http://localhost:8080/todos/'+ id);
			console.log("should be togeled")
		} catch (err) {
			console.log(err);
		}
	};  

  
    $: totalTodos = todos.length
    $: completedTodos = todos.filter((todo) => todo.Completed).length
    let newTodoItem = ''
    let newTodoId
    $: {
       if (totalTodos === 0) {
         newTodoId = 1;
        } 
        else {
          newTodoId = Math.max(...todos.map((t) => t.ID)) + 1;
        }
      }
    function removeTodo(todo) {
     todos = todos.filter((t) => t.ID !== todo.ID)
     remove_todo_api(todo.ID)
    }

    function remove_completed_todos () {
      for (let i = 0; i < todos.length; i++) {
         if (todos[i].Completed){
                removeTodo(todos[i])
                i--
           } 
    }
	};  

  function check_all_todos () {
      for (let i = 0; i < todos.length; i++) {
         if (!todos[i].Completed){
            todos[i].Completed = true
             update_todo_status_api(todos[i].ID)
           } 
    }
	};  
    function addTodo() {
      if (newTodoItem){
        todos = [...todos, { ID: newTodoId, Item: newTodoItem, Completed: false }]
        add_todo_api(newTodoItem)
        newTodoItem = '' 
        
      }
      else {
        alert("todo cannot be empty")
      }
    }

    let filter = 'all'
    const filterTodos = (filter, todos) =>
    filter === 'active' ? todos.filter((t) => !t.Completed) :
    filter === 'completed' ? todos.filter((t) => t.Completed) :
     todos

     $: editing = false
     $: edited_item = ""
     $: edit_todo_id = 0 
     $: edit_todo_item = "" 
    

     function on_edit(id , item){
       edit_todo_id = id
       edit_todo_item = item
       editing = true
     }

     function on_cancel(){
       console.log("why broo")
        editing = false
     }

     function on_save(id){
       update_todo_api( id , edited_item )
       for (let i = 0 ; i < todos.length ; i++){
          if (todos[i].ID === id){
            todos[i].Item = edited_item
            break
          }
       }
       edited_item = ""
       editing = false
     }
  </script>








<h1>Svelte to-do list</h1>
<title>Svelte to-do list</title>


<!-- Todos.svelte -->
<div class="todoapp stack-large">

    <!-- NewTodo -->
    <form on:submit|preventDefault={addTodo}>
      <h2 class="label-wrapper">
        <label for="todo-0" class="label__lg">
          What needs to be done?
        </label>
      </h2>
      <input bind:value={newTodoItem}  type="text" id="todo-0" autocomplete="off"
        class="input input__lg" />
      <button id="btn1" type="submit" disabled="" class="btn btn__primary btn__lg">
        Add
      </button>
    </form>
  
    <!-- Filter -->
    <div class="filters btn-group stack-exception">
      <button id="btn2" class="btn toggle-btn" class:btn__primary={filter === 'all'} 
              aria-pressed={filter === 'all'} on:click={() => filter = 'all'}>
        <span class="visually-hidden">Show</span>
        <span>All</span>
        <span class="visually-hidden">tasks</span>
      </button>
      <button id="btn3" class="btn toggle-btn" class:btn__primary={filter === 'active'} 
              aria-pressed={filter === 'active'}  on:click={() => filter = 'active'}>
        <span class="visually-hidden">Show</span>
        <span>Active</span>
        <span class="visually-hidden">tasks</span>
      </button>
      <button id="btn4" class="btn toggle-btn" class:btn__primary={filter === 'completed'} 
      aria-pressed={filter === 'completed'} on:click={() => filter = 'completed'}>
        <span class="visually-hidden">Show</span>
        <span>Completed</span>
        <span class="visually-hidden">tasks</span>
      </button>
    </div>
  
    <!-- TodosStatus -->
    <h2 id="list-heading">{completedTodos} out of {totalTodos} items completed</h2>
  
  <!-- To-dos -->
<ul role="list" id= "list1" class="todo-list stack-large" aria-labelledby="list-heading">
  {#if editing}
  <li class="todo">
    <div class="stack-small">  
      <form on:submit|preventDefault={on_save} class="stack-small" on:keydown={(e) => e.key === 'Escape' && on_cancel()}>
        <div class="form-group">
          <label for="todo-{edit_todo_id}" class="todo-label">New item for '{edit_todo_item}'</label>
          <input bind:value={edited_item} type="text" id="todo-edit-{edit_todo_id}" autoComplete="off" class="todo-text" />
        </div>
        <div class="btn-group">
          <button id="btn5-{edit_todo_id}"  class="btn todo-cancel" on:click={on_cancel} type="button">
            Cancel<span class="visually-hidden">renaming {edit_todo_item}</span>
            </button>
          <button id="btn6-{edit_todo_id}" class="btn btn__primary todo-edit" on:click={on_save(edit_todo_id)} type="submit" disabled={!edited_item}>
            Save<span class="visually-hidden">new item for {edit_todo_item}</span>
          </button>
        </div>
      </form>
    </div>
  </li>
  {:else} 
    {#each filterTodos(filter , todos) as todo (todo.ID)} 
    <li class="todo" id= "list2"  >
      <div class="stack-small">  
          <div class="c-cb">
            <input  type="checkbox"  on:click = {update_todo_status_api(todo.ID)} id="todo-check-{todo.ID}" 
            on:click={() => todo.Completed = !todo.Completed}
            checked={todo.Completed}/>
            <label for="todo-{todo.ID}" id="todo-text-{todo.ID}" class="todo-label">
              {todo.Item}
            </label>
          </div>
          <div class="btn-group">
            <button id="btn7-{todo.ID}"  type="button" class="btn" on:click={on_edit(todo.ID , todo.Item)}>
              Edit <span class="visually-hidden">{todo.Item}</span>
            </button>
            <button id="btn8-{todo.ID}" type="button" class="btn btn__danger"
             on:click={() => removeTodo(todo)}>
              Delete <span class="visually-hidden">{todo.Item}</span>
            </button>
            </div>  
          </div>
        </li>
        {:else}
        <li>There are no todos!</li>
      {/each}
    {/if}
       
        
   
  </ul>
    <!-- MoreActions -->
  <div class="btn-group">
    <button id="btn9" type="button" class="btn btn__primary"
    on:click={() => check_all_todos()}   >Check all</button>
    <button id="btn10" type="button" class="btn btn__primary"
      on:click={() => remove_completed_todos()} >Remove completed</button>
  </div>

</div>



    <style> 
    /* RESETS */
    *,
    *::before,
    *::after {
      box-sizing: border-box;
    }
    *:focus {
      outline: 3px dashed #228bec;
      outline-offset: 0;
    }
    html {
      font: 62.5% / 1.15 sans-serif;
    }
    h1,
    h2 {
      margin-bottom: 0;
    }
    ul {
      list-style: none;
      padding: 0;
    }
    button {
      border: none;
      margin: 0;
      padding: 0;
      width: auto;
      overflow: visible;
      background: transparent;
      color: inherit;
      font: inherit;
      line-height: normal;
      -webkit-font-smoothing: inherit;
      -moz-osx-font-smoothing: inherit;
      -webkit-appearance: none;
    }
    button::-moz-focus-inner {
      border: 0;
    }
    button,
    input,
    optgroup,
    select,
    textarea {
      font-family: inherit;
      font-size: 100%;
      line-height: 1.15;
      margin: 0;
    }
    button,
    input {
      overflow: visible;
    }
    input[type="text"] {
      border-radius: 0;
    }
    body {
      width: 100%;
      max-width: 68rem;
      margin: 0 auto;
      font: 1.6rem/1.25 Arial, sans-serif;
      background-color: #f5f5f5;
      color: #4d4d4d;
    }
    @media screen and (min-width: 620px) {
      body {
        font-size: 1.9rem;
        line-height: 1.31579;
      }
    }
    /*END RESETS*/
    
    /* GLOBAL STYLES */
    .form-group > input[type="text"] {
      display: inline-block;
      margin-top: 0.4rem;
    }
    .btn {
      padding: 0.8rem 1rem 0.7rem;
      border: 0.2rem solid #4d4d4d;
      cursor: pointer;
      text-transform: capitalize;
    }
    .btn.toggle-btn {
      border-width: 1px;
      border-color: #d3d3d3;
    }
    .btn.toggle-btn[aria-pressed="true"] {
      text-decoration: underline;
      border-color: #4d4d4d;
    }
    .btn__danger {
      color: #fff;
      background-color: #ca3c3c;
      border-color: #bd2130;
    }
    .btn__filter {
      border-color: lightgrey;
    }
    .btn__primary {
      color: #fff;
      background-color: #000;
    }
    .btn__primary:disabled {
      color: darkgrey;
      background-color:#565656;
    }
    .btn-group {
      display: flex;
      justify-content: space-between;
    }
    .btn-group > * {
      flex: 1 1 49%;
    }
    .btn-group > * + * {
      margin-left: 0.8rem;
    }
    .label-wrapper {
      margin: 0;
      flex: 0 0 100%;
      text-align: center;
    }
    .visually-hidden {
      position: absolute !important;
      height: 1px;
      width: 1px;
      overflow: hidden;
      clip: rect(1px 1px 1px 1px);
      clip: rect(1px, 1px, 1px, 1px);
      white-space: nowrap;
    }
    [class*="stack"] > * {
      margin-top: 0;
      margin-bottom: 0;
    }
    .stack-small > * + * {
      margin-top: 1.25rem;
    }
    .stack-large > * + * {
      margin-top: 2.5rem;
    }
    @media screen and (min-width: 550px) {
      .stack-small > * + * {
        margin-top: 1.4rem;
      }
      .stack-large > * + * {
        margin-top: 2.8rem;
      }
    }
    .stack-exception {
      margin-top: 1.2rem;
    }
    /* END GLOBAL STYLES */
    
    .todoapp {
      background: #fff;
      margin: 2rem 0 4rem 0;
      padding: 1rem;
      position: relative;
      box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2), 0 2.5rem 5rem 0 rgba(0, 0, 0, 0.1);
    }
    @media screen and (min-width: 550px) {
      .todoapp {
        padding: 4rem;
      }
    }
    .todoapp > * {
      max-width: 50rem;
      margin-left: auto;
      margin-right: auto;
    }
    .todoapp > form {
      max-width: 100%;
    }
    .todoapp > h1 {
      display: block;
      max-width: 100%;
      text-align: center;
      margin: 0;
      margin-bottom: 1rem;
    }
    .label__lg {
      line-height: 1.01567;
      font-weight: 300;
      padding: 0.8rem;
      margin-bottom: 1rem;
      text-align: center;
    }
    .input__lg {
      padding: 2rem;
      border: 2px solid #000;
    }
    .input__lg:focus {
      border-color: #4d4d4d;
      box-shadow: inset 0 0 0 2px;
    }
    [class*="__lg"] {
      display: inline-block;
      width: 100%;
      font-size: 1.9rem;
    }
    [class*="__lg"]:not(:last-child) {
      margin-bottom: 1rem;
    }
    @media screen and (min-width: 620px) {
      [class*="__lg"] {
        font-size: 2.4rem;
      }
    }
    .filters {
      width: 100%;
      margin: unset auto;
    }
    /* Todo item styles */
    .todo {
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;
    }
    .todo > * {
      flex: 0 0 100%;
    }
    .todo-text {
      width: 100%;
      min-height: 4.4rem;
      padding: 0.4rem 0.8rem;
      border: 2px solid #565656;
    }
    .todo-text:focus {
      box-shadow: inset 0 0 0 2px;
    }
    /* CHECKBOX STYLES */
    .c-cb {
      box-sizing: border-box;
      font-family: Arial, sans-serif;
      -webkit-font-smoothing: antialiased;
      font-weight: 400;
      font-size: 1.6rem;
      line-height: 1.25;
      display: block;
      position: relative;
      min-height: 44px;
      padding-left: 40px;
      clear: left;
    }
    .c-cb > label::before,
    .c-cb > input[type="checkbox"] {
      box-sizing: border-box;
      top: -2px;
      left: -2px;
      width: 44px;
      height: 44px;
    }
    .c-cb > input[type="checkbox"] {
      -webkit-font-smoothing: antialiased;
      cursor: pointer;
      position: absolute;
      z-index: 1;
      margin: 0;
      opacity: 0;
    }
    .c-cb > label {
      font-size: inherit;
      font-family: inherit;
      line-height: inherit;
      display: inline-block;
      margin-bottom: 0;
      padding: 8px 15px 5px;
      cursor: pointer;
      touch-action: manipulation;
    }
    .c-cb > label::before {
      content: "";
      position: absolute;
      border: 2px solid currentcolor;
      background: transparent;
    }
    .c-cb > input[type="checkbox"]:focus + label::before {
      border-width: 4px;
      outline: 3px dashed #228bec;
    }
    .c-cb > label::after {
      box-sizing: content-box;
      content: "";
      position: absolute;
      top: 11px;
      left: 9px;
      width: 18px;
      height: 7px;
      transform: rotate(-45deg);
      border: solid;
      border-width: 0 0 5px 5px;
      border-top-color: transparent;
      opacity: 0;
      background: transparent;
    }
    .c-cb > input[type="checkbox"]:checked + label::after {
      opacity: 1;
    }
    </style> 
  