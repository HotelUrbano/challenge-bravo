const MongoDBConnection = require('../config/mongo-db-config');

const COLLECTION_NAME = 'conversion-rates';

class ConversionRepository {
    constructor() {
        this.db = MongoDBConnection.db;
    }

    async insert(conversionRate) {
        await this.db.collection(COLLECTION_NAME).insertOne(conversionRate);
    }

    async getLatest() {
        const sorting = { referenceDate: -1 };
        const values = await this.db
            .collection(COLLECTION_NAME)
            .find()
            .sort(sorting)
            .limit(1)
            .toArray();
        return values[0];
    }
}

module.exports = new ConversionRepository();
