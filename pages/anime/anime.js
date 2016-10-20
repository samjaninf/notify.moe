const genreIcons = {
	'Action': 'bomb',
	'Adventure': 'diamond',
	'Cars': 'car',
	'Comedy': 'smile-o',
	'Drama': 'heartbeat',
	'Ecchi': 'heart-o',
	'Fantasy' : 'tree',
	'Game': 'gamepad',
	'Harem' : 'group',
	'Hentai': 'venus-mars',
	'Historical' : 'history',
	'Horror' : 'frown-o',
	'Kids': 'child',
	'Martial Arts' : 'hand-rock-o',
	'Magic': 'magic',
	'Mecha' : 'reddit-alien',
	'Military' : 'fighter-jet',
	'Music': 'music',
	'Mystery' : 'question',
	'Psychological': 'lightbulb-o',
	'Romance': 'heart',
	'Sci-Fi' : 'space-shuttle',
	'School' : 'graduation-cap',
	'Seinen' : 'male',
	'Shounen' : 'male',
	'Shoujo': 'female',
	'Slice of Life' : 'hand-peace-o',
	'Sports': 'soccer-ball-o',
	'Supernatural': 'magic',
	'Super Power': 'flash',
	'Thriller': 'hourglass-end',
	'Vampire': 'eye'
}

exports.get = function*(request, response) {
	let user = request.user
	let animeId = parseInt(request.params[0])
	let category = request.params[1]
	let categoryParameter = request.params[2]

	if(isNaN(animeId)) {
		let popular = yield arn.db.get('Cache', 'popularAnime')

		return response.render({
			user,
			popularAnime: popular.anime,
			animeToIdCount: arn.animeToIdCount,
			anime: null
		})
	}

	try {
		let animePage = yield arn.db.get('AnimePages', animeId)
		let videoParameters = ''

		if(category === 'video') {
			videoParameters += '&autoplay=1'

			if(categoryParameter && categoryParameter.indexOf('s') !== -1) {
				videoParameters += '&start=' + categoryParameter.replace('s', '')
			}
		}
		
		if(animePage.anime.studios) {
			animePage.anime.studios.forEach(studio => {
				studio.url = studio.wiki ? studio.wiki : `https://anilist.co/studio/${studio.id}`
			})
		}
		
		// Open Graph
		request.og = {
			url: app.package.homepage + '/anime/' + animePage.anime.id,
			title: animePage.anime.title.romaji,
			description: animePage.anime.description,
			image: animePage.anime.image
		}

		response.render(Object.assign({
			user,
			videoParameters,
			fixGenre: arn.fixGenre,
			nyaa: arn.animeProviders.Nyaa,
			genreIcons,
			canEdit: user && (user.role === 'admin' || user.role === 'editor'),
			friendsWatching: user ? animePage.usersWatching.filter(watcher => user.following.indexOf(watcher.id) !== -1) : null
		}, animePage))
	} catch(error) {
		console.error(error)

		response.render({
			user,
			error: 404
		})
	}
}