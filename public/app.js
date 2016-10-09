(function(Vue){
"use strict";

new Vue({
	el: 'body',

	data: {
		tasks: [],
		newTask: {}
	},

	created: function() {
		this.$http.get('/tasks').then(function(res) {
			this.tasks = res.data.items ? res.data.items : [];
		});
	},


	methods: {
		createTask: function() {
			if (!$.trim(this.newTask.name)) {
				this.newTask = {};
				return;
			};

			this.newTask.done = "false";

			this.$http.post('/task', this.newTask).success(function(res) {
				this.newTask.id = res.created;
				this.tasks.push(this.newTask);

				this.newTask = {};
			}).error(function(err) { console.log(err); });
		},


		deleteTask: function(index) {
			this.$http.delete('/task/' + index).success(function(res) {
				this.$http.get('/tasks').then(function(res) {
					this.tasks = res.data.items ? res.data.items : [];
				});
			}).error(function(err) { console.log(err); });
		},


		updateTask: function(task, completed) {
			if (completed) { task.done = "true"; }

			this.$http.put('/task', task).success(function(res) {
				this.$http.get('/tasks').then(function(res) {
					this.tasks = res.data.items ? res.data.items : [];
				});
			}).error(function(err) { console.log(err); });
		}
	}
});


})(Vue);
