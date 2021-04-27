# Canonical representation of Verses

## Goal

The idea here is that we define a human and machine readable format for storing verses.

### Format

The basic format is

```text
(?<bookid>[A-Za-z1-9]+) (?<chapter>[1-9][0-9]{0,2})/(?<verse_from>[1-9][0-9]{0,2})(-(?<verse_to>[1-9][0-9]{0,2}))?
```

Some tests can be found at [**https://regexr.com/5rmsd**](https://regexr.com/5rmsd)\*\*\*\*

Some valid examples:

* Exo 1/17
* 1Sam 2/18-34

### **Canonical book representation**

In order to be able to recognize the books, here is a reference list how they should be called

| Long name | Canonical Representation |
| :--- | :--- |
| Genesis | Gen |
| Exodus | Exo |
| Leviticus | Lev |
| Numbers | Num |
| Deuteronomy | Deu |
| Joshua | Josh |
| Judges | Judg |
| Ruth | Ruth |
| 1 Samuel | 1Sam |
| 2 Samuel | 2Sam |
| 1 Kings | 1Kgs |
| 2 Kings | 2Kgs |
| 1 Chronicles | 1Chr |
| 2 Chronicles | 2Chr |
| Ezra | Ezra |
| Nehemiah | Neh |
| 1 Esdras | 1Esd |
| Tobit | Tob |
| Judith | Jdt |
| Esther | Esth |
| Greek Esther | GkEst |
| Job | Job |
| Psalm | Ps |
| Psalm 151 | Ps151 |
| Proverbs | Prov |
| Ecclesiastes | Eccl |
| Song of Solomon | Song |
| Wisdom | Wis |
| Sirach | Sir |
| Isaiah | Isa |
| Jeremiah | Jer |
| Prayer of Azariah | PrAz |
| Lamentations | Lam |
| Letter of Jeremiah | EpJer |
| Baruch | Bar |
| Susanna | Sus |
| Ezekiel | Ezek |
| Daniel | Dan |
| Bel and the Dragon | Bel |
| Hosea | Hos |
| Joel | Joel |
| Amos | Am |
| Obadiah | Ob |
| Jonah | Jon |
| Micah | Mic |
| Nahum | Nah |
| Habakkuk | Hab |
| Zephaniah | Zeph |
| Haggai | Hag |
| Zechariah | Zech |
| Malachi | Mal |
| 1 Maccabees | 1Mac |
| 2 Maccabees | 2Mac |
| 3 Maccabees | 3Macc |
| 4 Maccabees | 4Macc |
| 2 Esdras | 2Esd |
| Matthew | Mt |
| Mark | Mk |
| Luke | Lk |
| John | Jn |
| Acts | Acts |
| Romans | Rom |
| 1 Corinthians | 1Cor |
| 2 Corinthians | 2Cor |
| Galatians | Gal |
| Ephesians | Eph |
| Philippians | Phil |
| Colossians | Col |
| 1 Thessalonians | 1Ths |
| 2 Thessalonians | 2Ths |
| 1 Timothy | 1Tim |
| 2 Timothy | 2Tim |
| Titus | Titus |
| Philemon | Phlm |
| Hebrews | Heb |
| James | Jas |
| 1 Peter | 1Pet |
| 2 Peter | 2Pet |
| 1 John | 1Jn |
| 2 John | 2Jn |
| 3 John | 3Jn |
| Jude | Jude |
| Revelation | Rev |
| Prayer of Manasseh | PrMan |

