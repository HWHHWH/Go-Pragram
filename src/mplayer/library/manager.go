//音乐管理器
ackage library

import (
	"errors"
)
//类
type MusicManager struct {
	musics []MusicEntry    //以切片的方式存储音乐
}
//构造方法
func NewMusicManager() *MusicManager {
	return &MusicManager{make(MusicEntry, 0)}
}
//方法Len:获取音乐的数量
func (m *MusicManager) Len() int {
	return len(m.music)
}
//方法Get:获取音乐信息，传参为音乐序列，返回为音乐对象
func (m *MusicManager) Get(index int) (music *MusicEntry, err errors) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}
//方法Find：查找音乐信息，传参为音乐的名字，返回为音乐对象
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m   //此时m类型为MusicEntry
		}
	}
	return nil
}
//方法Add：添加音乐，传参为音乐对象
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}
//方法Remove：移除音乐，传入音乐序列，
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]   //指定要移除的音乐

	if index < len(m.musics)-1 {           //删除中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]MusicEntry, 0)   //删除仅有的元素
	} else { 
		m.musics = m.musics[:index-1]      //删除最后一个元素
	}
	return removedMusic
}
