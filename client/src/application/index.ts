import { LayoutController } from 'controllers/layoutController'
import { LayoutStore } from 'stores/layoutStore'
import { RestApiService } from '../api'

export default class Application {
    apiService = new RestApiService()

    layoutStore = new LayoutStore()

    layoutController = new LayoutController(this.layoutStore)
}