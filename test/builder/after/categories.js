import { newOrLoadCase } from './layouts.js'
import { addCaseLoadButtonHandler, addNewCaseButtonHandler } from './cases.js'

export function writeCategorySelectList() {
    fetch('/after/categories').then(r => r.json()).then(categories => {
        const e = document.querySelector('#categories')
        for (const c of categories) {
            const o = document.createElement('option')
            o.value = c
            o.text = c
            e.add(o)
        }
    })
}

export function addCategoryLoadButtonHandler() {
    document.querySelector('#loadcategory').onclick = () => {
        const cat = document.querySelector('#categories')
        loadCategory(cat.options[cat.selectedIndex].value)
    }
}

let categories
export let pickedCategory

function loadCategory(c) {
    pickedCategory = c
    fetch('/after/category?name='+c).then(r => r.json()).then(cs => {
        categories = cs.cases
        document.querySelector('#content').innerHTML = newOrLoadCase
        const e = document.querySelector('#cases')
        for (const cat of categories) {
            const o = document.createElement('option')
            o.value = cat.case
            o.text = cat.case
            e.add(o)
        }
        addCaseLoadButtonHandler()
        addNewCaseButtonHandler()
    })
}

// If there is no loaded category matching name then an empty case is returned by categoryCase.
export function categoryCase(name) {
    for (const cas of categories) {
        if (cas.case !== name) {
            continue
        }
        return cas
    }
    return {
        active: 0,
        case: name,
        moves: [],
        pos: [],
        prev: {
            f: {
                f: 0,
                r: 0
            },
            t: {
                f: 0,
                r: 0
            }
        },
        state: 0
    }
}
