const neo4j = require('neo4j-driver').v1;
const async = require("async");
const driver = neo4j.driver('bolt://localhost:7687', neo4j.auth.basic('neo4j', 'class'));
const asyncForEach = require('async-foreach').forEach;
const session = driver.session()


async function getCount(){
    const result = await session.run('MATCH (n) RETURN count(n)')
    return result.records[0].get(0)
}

async function getNodes(){
    const res = await session.run('match (p:Person) return p limit 1')
    return res.records[0].get(0).properties.name
}

async function getEndorse(name){    
    const res = await session.run(`MATCH (:Person {name: '${name}'})-[:ENDORSES]->(other) RETURN distinct other`)
    return res
}

async function getEndorse2nd(name) {
    const res = await session.run(`MATCH (:Person {name: '${name}' })-[:ENDORSES]->()-[:ENDORSES]->(other_other) RETURN distinct other_other`)
    return res
}

async function getEndorse3rd(name) {
    const res = await session.run(`MATCH (:Person {name: '${name}' })-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->(other) RETURN distinct other`)
    return res
}

async function getEndorse4th(name) {
    const res = await session.run(`MATCH (:Person {name: '${name}' })-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->(other) RETURN distinct other`)
    return res
}

async function getEndorse5th(name) {
    const res = await session.run(`MATCH (:Person {name: '${name}' })-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->(other) RETURN distinct other`)
    return res
}

async function test(){
    const name = await getNodes()    

    console.time('Depth 1')
    const endorsements = await getEndorse(name)
    console.timeEnd('Depth 1')
    
    console.time('Depth 2')    
    const endorsements2nd = await getEndorse2nd(name)
    console.timeEnd('Depth 2')
    
    console.time('Depth 3')
    const endorsements3rd = await getEndorse3rd(name)
    console.timeEnd('Depth 3')
    
    console.time('Depth 4')    
    const endorsements4th = await getEndorse4th(name)
    console.timeEnd('Depth 4')
    
    console.time('Depth 5')
    const endorsements5th = await getEndorse5th(name)
    console.timeEnd('Depth 5')
    
    session.close()
    driver.close()
}

test()
