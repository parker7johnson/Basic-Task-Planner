


//create a function that will insert a new task into the dom
function insertNewTask(task) {
    //create a new list item
    let newTask = document.createElement('li');
    //add the task to the list item
    newTask.textContent = task;
    //append the new list item to the task list
    document.getElementById('task-list').appendChild(newTask);
}