const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const MoedaSchema = new Schema({
    moeda:{
        type: String,
        required: true,
        },
    cotacao: {
        type: Number,
        trim:true,
        required: false,
    },    
});

module.exports = mongoose.model('Moeda', MoedaSchema);

