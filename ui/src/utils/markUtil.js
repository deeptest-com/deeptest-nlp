export const convertSelectedToSlots = function (target, editor) {
  const allSlots = []
  const selectedSlot = {}

  target = getParentSpanNodeIfNeeded(target)
  console.log('---', target)

  const slt = window.getSelection()
  if (slt.toString() === '' || slt.rangeCount !== 1) return {}

  console.log('anchorNode', slt.anchorNode)
  console.log('anchorOffset', slt.anchorOffset)

  console.log('focusNode', slt.focusNode)
  console.log('focusOffset', slt.focusOffset)

  const range = slt.getRangeAt(0)
  console.log('range', range, range.toString())

  const startContainer = getParentSpanNodeIfNeeded(range.startContainer)
  const endContainer = getParentSpanNodeIfNeeded(range.endContainer)
  console.log('startContainer', startContainer, range.startOffset)
  console.log('endContainer', endContainer, range.endOffset)
  console.log('is same', startContainer === endContainer)

  let isStart = false
  for (let i = 0; i < editor.childNodes.length; i++) {
    let item = getParentSpanNodeIfNeeded(editor.childNodes[i])
    if (item.nodeName === '#text') {
      const span = document.createElement('span')
      span.innerText = item.nodeValue
      item = span
    }

    item.setAttribute('id', i + '')
    allSlots.push(item)

    if (item === startContainer) {
      console.log('start')
      isStart = true
    }

    if (isStart) {

    }

    if (item === endContainer) {
      console.log('end')
      break
    }
  }

  const mp = { allSlots: allSlots, selectedSlot: selectedSlot }
  console.log(mp)
  return mp

  // const items = []
  // let item = startContainer
  // while (true) {
  //   item = getParentSpanNodeIfNeeded(item)
  //
  //   let tp = item.nodeName.toLowerCase()
  //   tp = tp.replace('#', '')
  //   let html = ''
  //   let text = ''
  //   if (tp === 'span') {
  //     html = item.outerHTML
  //     text = item.innerText
  //   } else {
  //     html = item.wholeText
  //     text = item.wholeText
  //   }
  //
  //   console.log(tp, html)
  //   items.push({ elemType: tp, html: html, text: text })
  //
  //   if (item.nextSibling) {
  //     item = item.nextSibling
  //   } else {
  //     item = item.parentNode.nextSibling
  //   }
  //   if (!item) {
  //     break
  //   }
  // }
  //
  // const startText = startContainer.textContent
  // const endText = endContainer.textContent
  //
  // const startLeft = slt.anchorOffset
  // let startRight = startText.length
  // let endLeft = 0
  // const endRight = slt.focusOffset
  //
  // if (startContainer === endContainer) {
  //   startRight = endRight
  //   endLeft = startLeft
  // }
  //
  // items[0].selected = startText.substr(startLeft, startRight - startLeft)
  // console.log('start', items[0].selected, startLeft, startRight - startLeft)
  //
  // console.log(items)
  //
  // const selectedSize = range.toString().length
  // let totalSize = 0
  // const temp = []
  // for (let i = 0; i < items.length; i++) {
  //   const item = items[i]
  //   temp.push(item)
  //   const selected = item.selected ? item.selected : item.text
  //
  //   totalSize += selected.length
  //   if (totalSize >= selectedSize) {
  //     break
  //   }
  // }
  //
  // temp[temp.length - 1].selected = endText.substr(endLeft, endRight - endLeft)
  // console.log('end', temp[temp.length - 1].selected, endLeft, endRight - endLeft)
  //
  // console.log(temp)
  // return temp
}

export const genSent = function (sentItems, slot) {
  const arr = []

  sentItems.forEach((item, index) => {
    const section = document.createElement('span')
    section.setAttribute('data-seq', index)
    section.setAttribute('data-type', slot.slotType)
    section.setAttribute('data-value', slot.value)
    section.innerText = item.text

    arr.push(section.outerHTML)
  })

  const html = arr.join('')
  return html
}

export const getParentSpanNodeIfNeeded = function (target) {
  if (target.parentNode && target.parentNode.nodeName.toLowerCase() === 'span') {
    target = target.parentNode
  }
  return target
}
