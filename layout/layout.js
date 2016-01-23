'use strict'

exports.render = function(request, render) {
	let user = request.user
	let nav = []

	nav.push({
		title: 'Dash',
		url: '',
		icon: 'dashboard'
	})

	if(user) {
		nav.push({
			title: 'Profile',
			url: '+' + user.nick,
			icon: 'user'
		})
	}

	nav.push({
		title: 'Users',
		url: 'users',
		icon: 'globe'
	})

	nav.push({
		title: 'Anime',
		url: 'anime',
		icon: 'film'
	})

	if(user) {
		nav.push({
			title: 'Settings',
			url: 'settings',
			icon: 'cog'
		})
	} else {
		nav.push({
			title: 'FAQ',
			url: 'faq',
			icon: 'question-sign'
		})
	}

	if(user) {
		nav.push({
			title: '',
			url: 'logout',
			icon: 'log-out',
			ajax: false,
			float: 'right',
			tooltip: 'Logout'
		})

		nav.push({
			title: '',
			url: 'https://www.paypal.me/blitzprog',
			icon: 'heart',
			float: 'right',
			tooltip: 'Donate'
		})

		nav.push({
			title: '',
			url: 'api',
			icon: 'fire',
			float: 'right',
			tooltip: 'API (for developers)'
		})

		nav.push({
			title: '',
			url: 'faq',
			icon: 'question-sign',
			float: 'right',
			tooltip: 'FAQ'
		})

		nav.push({
			title: '',
			url: 'changes',
			icon: 'refresh',
			float: 'right',
			tooltip: 'Changes'
		})

		nav.push({
			title: '',
			url: 'roadmap',
			icon: 'road',
			float: 'right',
			tooltip: 'Roadmap'
		})

		nav.push({
			title: '',
			url: 'feedback',
			icon: 'comment',
			float: 'right',
			tooltip: 'Feedback'
		})

		if(user.role === 'admin' || user.role === 'editor') {
			nav.push({
				title: '',
				url: 'admin',
				icon: 'wrench',
				float: 'right',
				tooltip: 'Admin Panel'
			})
		}
	}

	render({
		user,
		nav,
		maintenance: arn.maintenance
	})
}